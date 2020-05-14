package scheduler

//go:generate sh -c "cd ../../../ && make mockgen"

import (
	"github.com/atlarge-research/opendc-emulate-kubernetes/api/apatelet"
	"github.com/atlarge-research/opendc-emulate-kubernetes/api/scenario"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/any"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/events"
	"github.com/atlarge-research/opendc-emulate-kubernetes/services/apatelet/store"
	"github.com/atlarge-research/opendc-emulate-kubernetes/services/apatelet/store/mock_store"

	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestTaskHandlerSimple(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	a, _ := any.Marshal(scenario.Response_RESPONSE_ERROR)

	// Test task:
	task := apatelet.Task{
		NodeEventFlags: map[int32]*anypb.Any{
			events.NodeCreatePodResponse: a,
		},
	}

	// Set up expectations
	ms.EXPECT().SetNodeFlag(events.NodeCreatePodResponse, scenario.Response_RESPONSE_ERROR)

	var s store.Store = ms
	sched := Scheduler{&s}

	// Run code under test
	ech := make(chan error)

	sched.taskHandler(ech, &store.Task{
		RelativeTimestamp: 0,
		IsPod:             false,
		PodTask:           nil,
		NodeTask:          &task,
	})

	select {
	case <-ech:
		t.Fail()
	default:
	}
}

func TestTaskHandlerMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	m1, err := any.Marshal(scenario.Response_RESPONSE_ERROR)
	assert.NoError(t, err)

	m2, err := any.Marshal(42)
	assert.NoError(t, err)

	// Test task:
	task := apatelet.Task{
		NodeEventFlags: map[int32]*anypb.Any{
			events.NodeCreatePodResponse: m1,
			events.NodeAddedLatencyMsec:  m2,
		},
	}

	// Set up expectations
	ms.EXPECT().SetNodeFlag(events.NodeCreatePodResponse, scenario.Response_RESPONSE_ERROR)
	ms.EXPECT().SetNodeFlag(events.NodeAddedLatencyMsec, int64(42))

	ms.EXPECT().SetPodFlag("a", events.PodCreatePodResponse, scenario.Response_RESPONSE_ERROR)
	ms.EXPECT().SetPodFlag("b", events.PodStatus, scenario.PodStatus_POD_STATUS_FAILED)

	var s store.Store = ms
	sched := Scheduler{&s}

	// Run code under test
	ech := make(chan error)

	sched.taskHandler(ech, &store.Task{
		RelativeTimestamp: 0,
		IsPod:             false,
		PodTask:           nil,
		NodeTask:          &task,
	})

	select {
	case <-ech:
		t.Fail()
	default:
	}
}

func TestTaskHandlerNodeError(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	a := anypb.Any{
		TypeUrl: "invalid-any",
		Value:   nil,
	}

	// Test task:
	task := apatelet.Task{
		NodeEventFlags: map[int32]*anypb.Any{
			events.NodeCreatePodResponse: &a,
		},
	}

	// Set up expectations
	var s store.Store = ms
	sched := Scheduler{&s}

	// Run code under test
	ech := make(chan error, 1)

	sched.taskHandler(ech, &store.Task{
		RelativeTimestamp: 0,
		IsPod:             false,
		PodTask:           nil,
		NodeTask:          &task,
	})

	select {
	case err := <-ech:
		assert.Error(t, err)
	default:
		t.Fail()
	}
}

func TestRunner(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	a, _ := any.Marshal(scenario.Response_RESPONSE_ERROR)

	// Test task:
	task := apatelet.Task{
		NodeEventFlags: map[int32]*anypb.Any{
			events.NodeCreatePodResponse: a,
		},
	}

	// Expectations
	ms.EXPECT().PeekTask().Return(int64(0), nil)
	ms.EXPECT().PopTask().Return(&store.Task{
		RelativeTimestamp: 0,
		IsPod:             false,
		PodTask:           nil,
		NodeTask:          &task,
	}, nil)
	ms.EXPECT().SetNodeFlag(gomock.Any(), gomock.Any())

	var s store.Store = ms
	sched := Scheduler{&s}

	// Run code under test
	ech := make(chan error, 1)

	sched.runner(ech)

	select {
	case <-ech:
		t.Fail()
	default:
	}
}

func TestRunnerEarlyFail(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	// Expectations
	ms.EXPECT().PeekTask().Return(int64(-1), errors.New("some error"))

	// Setup
	var s store.Store = ms
	sched := Scheduler{&s}

	// Test
	ech := make(chan error, 1)
	sched.runner(ech)

	// Verify
	select {
	case err := <-ech:
		assert.Error(t, err)
	default:
		t.Fail()
	}
}

func TestRunnerPopFail(t *testing.T) {
	ctrl := gomock.NewController(t)

	ms := mock_store.NewMockStore(ctrl)

	// Expectations
	ms.EXPECT().PeekTask().Return(int64(0), nil)
	ms.EXPECT().PopTask().Return(nil, errors.New("new error"))

	// Setup
	var s store.Store = ms
	sched := Scheduler{&s}

	// Test
	ech := make(chan error, 1)
	sched.runner(ech)

	// Verify
	select {
	case err := <-ech:
		assert.Error(t, err)
	default:
		t.Fail()
	}
}

func TestStartScheduler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctrl, ctx := gomock.WithContext(ctx, t)

	ms := mock_store.NewMockStore(ctrl)

	a, _ := any.Marshal(scenario.Response_RESPONSE_ERROR)

	// Test task:
	task := apatelet.Task{
		NodeEventFlags: map[int32]*anypb.Any{
			events.NodeCreatePodResponse: a,
		},
	}

	// Expectations
	ms.EXPECT().PeekTask().Return(int64(0), nil)
	ms.EXPECT().PopTask().Return(&store.Task{
		RelativeTimestamp: 0,
		IsPod:             false,
		PodTask:           nil,
		NodeTask:          &task,
	}, nil)
	ms.EXPECT().SetNodeFlag(gomock.Any(), gomock.Any())

	// any further peeks are well into the future
	ms.EXPECT().PeekTask().Return(time.Now().Add(time.Hour*12).UnixNano(), nil).AnyTimes()

	var s store.Store = ms
	sched := Scheduler{&s}

	// Run code under test
	ech := sched.StartScheduler(ctx)

	time.Sleep(time.Second)

	select {
	case <-ech:
		t.Fail()
	default:
	}
}
