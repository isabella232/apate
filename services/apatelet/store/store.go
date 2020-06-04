// Package store provides a way for the apatelet to have state
package store

import (
	"container/heap"
	corev1 "k8s.io/api/core/v1"
	"sync"
	"time"

	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"

	"github.com/pkg/errors"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario"

	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/events"
)

// Store represents the state of the apatelet
type Store interface {
	// SetNodeTasks adds or updates node tasks
	// Existing node tasks will be removed if not in the list of tasks
	SetNodeTasks([]*Task) error

	// SetPodTasks adds or updates pod CRD tasks to the queue based on their label (<namespace>/<name>)
	// Existing pod tasks will be removed if not in the list of tasks
	SetPodTasks(string, []*Task) error

	// SetNodeFlag sets the value of the given pod flag for a configuration
	SetPodTimeFlags(string, []*TimeFlags)

	// RemovePodTasks removes pod CRD tasks from the queue based on their label (<namespace>/<name>)
	RemovePodTasks(string) error

	// LenTasks returns the amount of tasks left to be picked up
	LenTasks() int

	// PeekTask returns the start time of the next task in the priority queue, without removing it from the queue
	PeekTask() (time.Duration, bool, error)

	// PopTask returns the first task to be executed and removes it from the queue
	PopTask() (*Task, error)

	// GetNodeFlag returns the value of the given node flag
	GetNodeFlag(events.NodeEventFlag) (interface{}, error)

	// SetNodeFlag sets the value of the given node flag
	SetNodeFlag(events.NodeEventFlag, interface{})

	// GetPodFlag returns the value of the given pod flag for a configuration
	GetPodFlag(string, *corev1.Pod, events.PodEventFlag) (interface{}, error)

	// SetNodeFlag sets the value of the given pod flag for a configuration
	SetPodFlags(string, Flags)

	// AddPodListener adds a listener which is called when the given flag is updated
	AddPodListener(events.PodEventFlag, func(interface{}))
}

type Flags map[events.EventFlag]interface{}
type podFlags map[string]Flags
type podListeners map[events.EventFlag][]func(interface{})

type TimeFlags struct {
	TimeSincePodStart time.Duration
	Flags             Flags
}
type podTimeFlags map[string][]*TimeFlags
type podTimeIndexCache map[*corev1.Pod]map[events.EventFlag]int

type store struct {
	queue     *taskQueue
	queueLock sync.RWMutex

	nodeFlags    Flags
	nodeFlagLock sync.RWMutex

	podFlags    podFlags
	podFlagLock sync.RWMutex

	podListeners     podListeners
	podListenersLock sync.RWMutex

	podTimeFlags      podTimeFlags
	podTimeIndexCache podTimeIndexCache
}

// NewStore returns an empty store
func NewStore() Store {
	q := newTaskQueue()
	heap.Init(q)

	return &store{
		queue:        q,
		nodeFlags:    make(Flags),
		podListeners: make(podListeners),
		podFlags:     make(podFlags),

		podTimeFlags:      make(podTimeFlags),
		podTimeIndexCache: make(podTimeIndexCache),
	}
}

func (s *store) setTasksOfType(newTasks []*Task, check TaskTypeCheck) error {
	s.queueLock.Lock()
	defer s.queueLock.Unlock()

	for i, task := range s.queue.tasks {
		typeCheck, err := check(task)
		if err != nil {
			return errors.Wrap(err, "failed to determine task type")
		}

		if typeCheck {
			if len(newTasks) == 0 {
				heap.Remove(s.queue, i)
			} else {
				if newTasks[0] != nil {
					s.queue.tasks[i] = newTasks[0]
					// Replacing and then fixing instead of deleting all and pushing because it's slightly faster, see comments on heap.Fix
					heap.Fix(s.queue, i)
				}
				newTasks = newTasks[1:]
			}
		}
	}

	for _, remainingTask := range newTasks {
		if remainingTask != nil {
			heap.Push(s.queue, remainingTask)
		}
	}

	return nil
}

func (s *store) SetNodeTasks(tasks []*Task) error {
	return s.setTasksOfType(tasks, func(task *Task) (bool, error) {
		isNode, err := task.IsNode()
		if err != nil {
			return false, err
		}

		return isNode, nil
	})
}

func (s *store) SetPodTasks(label string, tasks []*Task) error {
	return s.setTasksOfType(tasks, func(task *Task) (bool, error) {
		isPod, err := task.IsPod()
		if err != nil {
			return false, err
		}

		return isPod && task.PodTask.Label == label, nil
	})
}

func (s *store) SetPodTimeFlags(label string, flags []*TimeFlags) {
	s.podTimeFlags[label] = flags
}

func (s *store) RemovePodTasks(label string) error {
	s.queueLock.Lock()
	defer s.queueLock.Unlock()

	for i := len(s.queue.tasks) - 1; i >= 0; i-- {
		task := s.queue.tasks[i]

		isPod, err := task.IsPod()
		if err != nil {
			return errors.Wrap(err, "failed to determine task type")
		}

		if isPod && task.PodTask.Label == label {
			heap.Remove(s.queue, i)
		}
	}

	return nil
}

func (s *store) LenTasks() int {
	s.queueLock.RLock()
	defer s.queueLock.RUnlock()

	return s.queue.Len()
}

func (s *store) PeekTask() (time.Duration, bool, error) {
	s.queueLock.RLock()
	defer s.queueLock.RUnlock()

	if s.queue.Len() == 0 {
		return -1, false, nil
	}

	// Make sure the array in the pq didn't magically change to a different type
	if task, ok := s.queue.First().(*Task); ok {
		return task.RelativeTimestamp, true, nil
	}

	return -1, false, errors.New("array in pq magically changed to a different type")
}

func (s *store) PopTask() (*Task, error) {
	s.queueLock.Lock()
	defer s.queueLock.Unlock()

	if s.queue.Len() == 0 {
		return nil, errors.New("no tasks left")
	}

	// Make sure the array in the pq didn't magically change to a different type
	if task, ok := heap.Pop(s.queue).(*Task); ok {
		return task, nil
	}

	return nil, errors.New("array in pq magically changed to a different type")
}

func (s *store) GetNodeFlag(id events.NodeEventFlag) (interface{}, error) {
	s.nodeFlagLock.RLock()
	defer s.nodeFlagLock.RUnlock()

	if val, ok := s.nodeFlags[id]; ok {
		return val, nil
	}

	if dv, ok := defaultNodeValues[id]; ok {
		return dv, nil
	}

	return nil, errors.New("flag not found in get node flag")
}

func (s *store) SetNodeFlag(id events.NodeEventFlag, val interface{}) {
	s.nodeFlagLock.Lock()
	defer s.nodeFlagLock.Unlock()

	s.nodeFlags[id] = val
}

func (s *store) GetPodFlag(label string, pod *corev1.Pod, flag events.PodEventFlag) (interface{}, error) {
	s.podFlagLock.Lock()
	defer s.podFlagLock.Unlock()

	if label != "" {
		if val, ok := s.podFlags[label][flag]; ok {
			return val, nil
		}
	}

	if _, ok := s.podTimeIndexCache[pod]; !ok {
		s.podTimeIndexCache[pod] = make(map[events.EventFlag]int)
	}

	podTimeIndex := 0
	if val, ok := s.podTimeIndexCache[pod][flag]; ok {
		podTimeIndex = val
	}

	podStartTime := time.Now()
	if pod.Status.StartTime != nil {
		podStartTime = pod.Status.StartTime.Time
	}

	timeFlags := s.podTimeFlags[label]
	lastIndexWithFlag := podTimeIndex
	for i := podTimeIndex; i < len(timeFlags); i++ {
		flags := timeFlags[i]

		if podStartTime.Add(flags.TimeSincePodStart).Before(time.Now()) {
			currentPodFlags := timeFlags[lastIndexWithFlag]
			s.podTimeIndexCache[pod][flag] = lastIndexWithFlag
			return currentPodFlags.Flags[flag], nil
		}

		if _, ok := flags.Flags[flag]; ok {
			lastIndexWithFlag = i
		}
	}

	if dv, ok := defaultPodValues[flag]; ok {
		return dv, nil
	}

	return nil, errors.New("flag not found in get pod flag")
}

func (s *store) SetPodFlags(label string, flags Flags) {
	s.podFlagLock.Lock()
	s.podFlags[label] = flags
	s.podFlagLock.Unlock()

	s.podListenersLock.RLock()
	for flag, val := range flags {
		if listeners, ok := s.podListeners[flag]; ok {
			for _, listener := range listeners {
				listener(val)
			}
		}
	}
	s.podListenersLock.RUnlock()
}

func (s *store) AddPodListener(flag events.PodEventFlag, cb func(interface{})) {
	s.podListenersLock.Lock()
	defer s.podListenersLock.Unlock()

	if listeners, ok := s.podListeners[flag]; ok {
		s.podListeners[flag] = append(listeners, cb)
	} else {
		s.podListeners[flag] = []func(interface{}){cb}
	}
}

var defaultNodeValues = map[events.EventFlag]interface{}{
	events.NodeCreatePodResponse:    scenario.ResponseUnset,
	events.NodeUpdatePodResponse:    scenario.ResponseUnset,
	events.NodeDeletePodResponse:    scenario.ResponseUnset,
	events.NodeGetPodResponse:       scenario.ResponseUnset,
	events.NodeGetPodStatusResponse: scenario.ResponseUnset,
	events.NodeGetPodsResponse:      scenario.ResponseUnset,
	events.NodePingResponse:         scenario.ResponseUnset,

	events.NodeAddedLatency: time.Duration(0),
}

var defaultPodValues = map[events.PodEventFlag]interface{}{
	events.PodCreatePodResponse:    scenario.ResponseUnset,
	events.PodUpdatePodResponse:    scenario.ResponseUnset,
	events.PodDeletePodResponse:    scenario.ResponseUnset,
	events.PodGetPodResponse:       scenario.ResponseUnset,
	events.PodGetPodStatusResponse: scenario.ResponseUnset,

	events.PodResources: &stats.PodStats{},

	events.PodStatus: scenario.PodStatusUnset,
}
