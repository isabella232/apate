// Package kubeconfig provides the ability to create, read, and manage the kubeconfig file/bytes.
package kubeconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeConfig is an alias of a bytearray, and represents a raw kube configuration file loaded from file.
type KubeConfig struct {
	Path  string
	Bytes []byte
}

// FromBytes creates a kubeConfig struct from byte array.
func FromBytes(bytes []byte) (*KubeConfig, error) {
	path := os.TempDir() + "/apate/config-" + uuid.New().String()

	if err := ioutil.WriteFile(path, bytes, 0o600); err != nil {
		return nil, err
	}

	return &KubeConfig{
		Path:  path,
		Bytes: bytes,
	}, nil
}

// FromPath Loads a KubeConfig from a file path.
func FromPath(path string) (*KubeConfig, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return &KubeConfig{path, bytes}, nil
}

// GetConfig returns a kubernetes rest configuration from the KubeConfig.
func (k KubeConfig) GetConfig() (*rest.Config, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig(k.Bytes)
	if err != nil {
		return nil, err
	}

	return config, nil
}