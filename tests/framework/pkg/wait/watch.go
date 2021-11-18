package wait

import (
	"k8s.io/apimachinery/pkg/watch"
)

func WatchWait(watchInterface watch.Interface, check CheckFunc) error {
	return nil
}
