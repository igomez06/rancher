package session

import (
	"testing"

	"github.com/rancher/rancher/tests/framework/pkg/wait"
	"k8s.io/apimachinery/pkg/watch"
)

type CleanupFunc func() error
type WaitFunc func(watchInterface watch.Interface, check wait.CheckFunc) error

type Session struct {
	cleanupQueue []CleanupFunc
	waitQueue    []CleanupFunc
	open         bool
	testingT     *testing.T
}

func NewSession(t *testing.T) *Session {
	return &Session{
		cleanupQueue: []CleanupFunc{},
		open:         true,
		testingT:     t,
	}
}

// RegisterCleanupFunc functions passed to this method will be called in the order they are added when `Cleanup` is called.
// If Session is closed, it will cause a panic if a new cleanup function is registered.
func (ts *Session) RegisterCleanupFunc(f CleanupFunc) {
	if ts.open {
		ts.cleanupQueue = append(ts.cleanupQueue, f)
	} else {
		panic("attempted to register cleanup function to closed test session")
	}

}

// RegisterWaitFunc
func (ts *Session) RegisterWaitFunc(f CleanupFunc) {
	if ts.open {
		ts.waitQueue = append(ts.waitQueue, f)
	} else {
		panic("attempted to register cleanup function to closed test session")
	}

}

// Cleanup this method will call all registered cleanup functions in order and close the test session.
func (ts *Session) Cleanup() {
	ts.open = false

	count := len(ts.cleanupQueue) - 1

	for index := range ts.cleanupQueue {
		itemIndex := count - index
		err := ts.cleanupQueue[itemIndex]()
		if err != nil {
			ts.testingT.Logf("error calling cleanup function: %v", err)
		}

		err = ts.waitQueue[itemIndex]()
		if err != nil {
			ts.testingT.Logf("error waiting on resource: %v", err)
		}
	}
}

// NewSession returns a `Session` who's cleanup method is registered with this `Session`
func (ts *Session) NewSession() *Session {
	sess := NewSession(ts.testingT)

	ts.RegisterCleanupFunc(func() error {
		sess.Cleanup()
		return nil
	})

	return sess
}
