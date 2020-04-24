// Package store provides state to the apate cluster
package store

import (
	"fmt"
	"sync"

	"github.com/atlarge-research/opendc-emulate-kubernetes/api/apatelet"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/normalization"

	"github.com/google/uuid"
)

//TODO: Multi-master soon :tm:

// Store represents the store of the control plane
type Store interface {
	// AddNode adds the given Node to the Apate cluster
	AddNode(*Node) error

	// RemoveNode removes the given Node from the Apate cluster
	RemoveNode(*Node) error

	// GetNode returns the node with the given uuid
	GetNode(uuid.UUID) (Node, error)

	// GetNodes returns an array containing all nodes in the Apate cluster
	GetNodes() ([]Node, error)

	// ClearNodes removes all nodes from the Apate cluster
	ClearNodes() error

	// AddResourcesToQueue adds a node resource to the queue
	AddResourcesToQueue([]normalization.NodeResources) error

	// AddApateletScenario adds the ApateletScenario to the store
	AddApateletScenario(*apatelet.ApateletScenario) error

	// GetApateletScenario gets the ApateletScenario
	GetApateletScenario() (*apatelet.ApateletScenario, error)
}

type store struct {
	nodes    map[uuid.UUID]Node
	nodeLock sync.RWMutex
}

// NewStore creates a new empty cluster
func NewStore() Store {
	return &store{
		nodes: make(map[uuid.UUID]Node),
	}
}

// AddNode adds the given Node to the Apate cluster
func (c *store) AddNode(node *Node) error {
	c.nodeLock.Lock()
	defer c.nodeLock.Unlock()

	// Check if node already exists (uuid collision)
	if _, ok := c.nodes[node.UUID]; ok {
		return fmt.Errorf("node with uuid '%s' already exists", node.UUID.String())
	}

	c.nodes[node.UUID] = *node

	return nil
}

// RemoveNode removes the given Node from the Apate cluster
func (c *store) RemoveNode(node *Node) error {
	c.nodeLock.Lock()
	defer c.nodeLock.Unlock()

	delete(c.nodes, node.UUID)
	return nil
}

// GetNode returns the node with the given uuid
func (c *store) GetNode(uuid uuid.UUID) (Node, error) {
	c.nodeLock.RLock()
	defer c.nodeLock.RUnlock()

	if node, ok := c.nodes[uuid]; ok {
		return node, nil
	}

	return Node{}, fmt.Errorf("node with uuid '%s' not found", uuid.String())
}

// GetNodes returns an array containing all nodes in the Apate cluster
func (c *store) GetNodes() ([]Node, error) {
	c.nodeLock.RLock()
	defer c.nodeLock.RUnlock()

	nodes := make([]Node, 0, len(c.nodes))

	for _, node := range c.nodes {
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (c *store) ClearNodes() error {
	c.nodeLock.Lock()
	defer c.nodeLock.Unlock()

	c.nodes = make(map[uuid.UUID]Node)
	return nil
}

func (c *store) AddResourcesToQueue(resources []normalization.NodeResources) error {
	return nil
}

func (c *store) AddApateletScenario(scenario *apatelet.ApateletScenario) error {
	return nil
}

func (c *store) GetApateletScenario() (*apatelet.ApateletScenario, error) {
	return nil, nil
}
