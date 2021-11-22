package wait

import (
	"fmt"

	"k8s.io/apimachinery/pkg/watch"
)

type WatchCheckFunc func(watch.Event) (bool, error)

func WatchWait(watchInterface watch.Interface, check WatchCheckFunc) error {
	defer func() {
		watchInterface.Stop()
	}()

	for {
		select {
		case event, open := <-watchInterface.ResultChan():
			if !open {
				return fmt.Errorf("timeout waiting on condition")
			}

			done, err := check(event)
			if err != nil {
				return err
			}

			if done {
				return nil
			}
		}
	}
}
