package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/kubectl"

	emulatedpodv1 "github.com/atlarge-research/opendc-emulate-kubernetes/pkg/apis/emulatedpod/v1"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/network"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/env"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/cluster"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/service"
	"github.com/atlarge-research/opendc-emulate-kubernetes/services/controlplane/services"
	"github.com/atlarge-research/opendc-emulate-kubernetes/services/controlplane/store"
)

func init() {
	// Enable line numbers in logging
	// Enables date time flags & file name + line
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("starting Apate control plane")

	// Get external connection information
	externalInformation, err := createExternalConnectionInformation()

	if err != nil {
		log.Fatalf("error while starting control plane: %s", err.Error())
	}

	// Create kubernetes cluster
	log.Println("starting kubernetes control plane")
	managedKubernetesCluster := createCluster(env.RetrieveFromEnvironment(env.ManagedClusterConfig, env.ManagedClusterConfigDefault))

	// Create apate cluster state
	createdStore := store.NewStore()

	// Save the kubeconfig in the store
	if err = createdStore.SetKubeConfig(*managedKubernetesCluster.KubeConfig); err != nil {
		log.Fatal(err)
	}

	if err = emulatedpodv1.CreateInKubernetes(managedKubernetesCluster.KubeConfig); err != nil {
		log.Fatal(err)
	}

	// Create prometheus stack
	createPrometheus := env.RetrieveFromEnvironment(env.PrometheusStackEnabled, env.PrometheusStackEnabledDefault)
	if strings.ToLower(createPrometheus) == "true" {
		go kubectl.CreatePrometheusStack(managedKubernetesCluster.KubeConfig)
	}

	// Start gRPC server
	server, err := createGRPC(&createdStore, managedKubernetesCluster.KubernetesCluster, externalInformation)
	if err != nil {
		log.Fatalf("error while starting control plane: %s", err.Error())
	}

	log.Printf("now accepting requests on %s:%d\n", server.Conn.Address, server.Conn.Port)

	if err = ioutil.WriteFile(os.TempDir()+"/apate/config", managedKubernetesCluster.KubernetesCluster.KubeConfig.Bytes, 0600); err != nil {
		log.Fatalf("error while starting control plane: %s", err.Error())
	}

	// Handle signals
	signals := make(chan os.Signal, 1)
	stopped := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		shutdown(&createdStore, &managedKubernetesCluster, server)
		stopped <- true
	}()

	// Start serving request
	go server.Serve()

	// Stop the server on signal
	<-stopped
	log.Printf("apate control plane stopped")
}

func shutdown(store *store.Store, kubernetesCluster *cluster.ManagedCluster, server *service.GRPCServer) {
	log.Println("stopping Apate control plane")

	log.Println("stopping API")
	server.Server.Stop()

	// TODO: Actual cleanup for other nodes, for now just wipe state
	if err := (*store).ClearNodes(); err != nil {
		log.Printf("an error occurred while cleaning the apate store: %s", err.Error())
	}

	log.Println("stopping kubernetes control plane")
	if err := kubernetesCluster.Delete(); err != nil {
		log.Printf("an error occurred while deleting the kubernetes store: %s", err.Error())
	}
}

func getExternalAddress() (string, error) {
	// Check for external IP override
	override := env.RetrieveFromEnvironment(env.ControlPlaneExternalIP, env.ControlPlaneExternalIPDefault)
	if override != env.ControlPlaneExternalIPDefault {
		return override, nil
	}

	return network.GetExternalAddress()
}

func createGRPC(createdStore *store.Store, kubernetesCluster cluster.KubernetesCluster, info *service.ConnectionInfo) (*service.GRPCServer, error) {
	// Retrieve from environment
	listenAddress := env.RetrieveFromEnvironment(env.ControlPlaneListenAddress, env.ControlPlaneListenAddressDefault)

	// Connection settings
	connectionInfo := service.NewConnectionInfo(listenAddress, info.Port, false)

	// Create gRPC server
	server, err := service.NewGRPCServer(connectionInfo)
	if err != nil {
		return nil, err
	}

	// Add services
	services.RegisterStatusService(server, createdStore)
	services.RegisterScenarioService(server, createdStore, info)
	services.RegisterClusterOperationService(server, createdStore, kubernetesCluster)
	services.RegisterHealthService(server, createdStore)

	return server, nil
}

func createCluster(managedClusterConfigPath string) cluster.ManagedCluster {
	cb := cluster.Default()
	c, err := cb.WithName("Apate").WithManagerConfig(managedClusterConfigPath).ForceCreate()
	if err != nil {
		log.Fatalf("an error occurred: %s", err.Error())
	}

	numberOfPods, err := c.GetNumberOfPods("kube-system")
	if err != nil {
		log.Fatalf("an error occurred: %s", err.Error())
	}

	log.Printf("There are %d pods in the cluster", numberOfPods)

	return c
}

func createExternalConnectionInformation() (*service.ConnectionInfo, error) {
	// Get external ip
	externalIP, err := getExternalAddress()

	if err != nil {
		return nil, err
	}

	// Get port
	listenPort, err := strconv.Atoi(env.RetrieveFromEnvironment(env.ControlPlaneListenPort, env.ControlPlaneListenPortDefault))

	if err != nil {
		return nil, err
	}

	// Create external information
	log.Printf("external IP for control plane: %s", externalIP)

	return service.NewConnectionInfo(externalIP, listenPort, false), nil
}
