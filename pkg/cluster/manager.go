// Defines an interface used to set up a cluster.
package cluster

type Manager interface {

	// Deletes a cluster with a given name.
	// This should never error, and should do nothing if the cluster to be deleted did not exist.
	DeleteCluster(name string) error

	// Should delete a cluster with a certain name.
	// This may error, and should error when a cluster with that name already exists.
	CreateCluster(name string, kubeConfigLocation string) error

	// Returns the name of a context for kubernetes to use for a given cluster name.
	ClusterContext(name string) string
}