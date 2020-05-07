package run

import (
	"context"
	healthpb "github.com/atlarge-research/opendc-emulate-kubernetes/api/health"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/clients/controlplane"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/clients/health"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/cluster"
	ec "github.com/atlarge-research/opendc-emulate-kubernetes/pkg/env"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/normalization"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/service"
	vkProvider "github.com/atlarge-research/opendc-emulate-kubernetes/services/apatelet/provider"
	vkService "github.com/atlarge-research/opendc-emulate-kubernetes/services/apatelet/services"
	"github.com/atlarge-research/opendc-emulate-kubernetes/services/apatelet/store"
	cli "github.com/virtual-kubelet/node-cli"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var KubeConfigWriter = func(config []byte) {
	// Noop by default
}

func StartApatelet(env ec.ApateletEnvironment, kubernetesPort, metricsPort int) error {
	SetCerts()
	log.Println("Starting Apatelet")

	// Retrieving connection information
	connectionInfo := service.NewConnectionInfo(env.ControlPlaneAddress, env.ControlPlanePort, false)
	ctx := context.Background()

	// Join the apate cluster
	log.Println("Joining apate cluster")
	config, res := joinApateCluster(ctx, connectionInfo, env.ListenPort)
	KubeConfigWriter(config)

	k, err := cluster.KubernetesClusterFromKubeConfig(config)

	if err != nil {
		return err
	}

	n, err := k.GetNumberOfPods("kube-system")
	if err != nil {
		return err
	}
	log.Printf("Pods: %d\n", n)

	// Setup health status
	hc := health.GetClient(connectionInfo, res.UUID.String())
	hc.SetStatus(healthpb.Status_UNKNOWN)
	hc.StartStreamWithRetry(ctx, 3)

	// Start the Apatelet
	nc, cancel, err := createNodeController(ctx, res, kubernetesPort, metricsPort)

	// Create virtual kubelet
	errch := make(chan error)

	log.Println("Joining kubernetes cluster")
	go func() {
		// TODO: Notify master / proper logging
		if err = nc.Run(); err != nil {
			hc.SetStatus(healthpb.Status_UNHEALTHY)
			errch <- err
		}
	}()

	st := store.NewStore()

	// Start gRPC server
	server, err := createGRPC(env.ListenPort, &st, env.ListenAddress)
	if err != nil {
		return err
	}

	// Update status
	hc.SetStatus(healthpb.Status_HEALTHY)
	log.Printf("Now accepting requests on %s:%d\n", server.Conn.Address, server.Conn.Port)

	// Handle signals
	signals := make(chan os.Signal, 1)
	stopped := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//TODO: Bubble up error to run.go
		err := <-errch
		log.Print(err)
		signals <- os.Kill
	}()

	go func() {
		<-signals
		shutdown(ctx, server, cancel, connectionInfo, res.UUID.String())
		stopped <- true
	}()

	// Start serving request
	server.Serve()

	// Stop the server on signal
	<-stopped
	log.Println("Apatelet stopped")
	return nil
}

func shutdown(ctx context.Context, server *service.GRPCServer, cancel context.CancelFunc, connectionInfo *service.ConnectionInfo, uuid string) {
	log.Println("Stopping Apatelet")

	log.Println("Stopping API")
	server.Server.Stop()

	log.Println("Leaving clusters (apate & k8s)")

	client := controlplane.GetClusterOperationClient(connectionInfo)
	defer func() {
		_ = client.Conn.Close()
	}()

	if err := client.LeaveCluster(ctx, uuid); err != nil {
		log.Printf("An error occurred while leaving the clusters (apate & k8s): %s", err.Error())
	}

	log.Println("Stopping provider")
	cancel()
}

func joinApateCluster(ctx context.Context, connectionInfo *service.ConnectionInfo, listenPort int) (cluster.KubeConfig, *normalization.NodeResources) {
	client := controlplane.GetClusterOperationClient(connectionInfo)
	defer func() {
		_ = client.Conn.Close()
	}()

	kubeconfig, res, err := client.JoinCluster(ctx, listenPort)

	// TODO: Better error handling
	if err != nil {
		log.Fatalf("Unable to join cluster: %v", err)
	}

	log.Printf("Joined apate cluster with resources: %v", res)

	return kubeconfig, res
}

func createNodeController(ctx context.Context, res *normalization.NodeResources, k8sPort int, metricsPort int) (*cli.Command, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(ctx)
	cmd, err := vkProvider.CreateProvider(ctx, res, k8sPort, metricsPort)
	return cmd, cancel, err
}

func createGRPC(listenPort int, store *store.Store, listenAddress string) (*service.GRPCServer, error) {

	// Connection settings
	connectionInfo := service.NewConnectionInfo(listenAddress, listenPort, false)

	// Create gRPC server
	server, err := service.NewGRPCServer(connectionInfo)
	if err != nil {
		return nil, err
	}

	// Add services
	vkService.RegisterScenarioService(server, store)

	return server, nil
}
