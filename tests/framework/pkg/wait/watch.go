package wait

import (
	"fmt"

	"k8s.io/apimachinery/pkg/watch"
)

func WatchWait(watchInterface watch.Interface, check CheckFunc) error {
	defer func() {
		watchInterface.Stop()
	}()

	for {
		select {
		case event, open := <-watchInterface.ResultChan():
			if !open {
				return fmt.Errorf("timeout waiting on condition")
			}
			switch event.Type {
			case watch.Modified:
				err := wait(check)
				if err != nil {
					return err
				}
				return nil
			case watch.Deleted:
				err := wait(check)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
}
