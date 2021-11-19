package wait

import (
	"fmt"
	"time"
)

type CheckFunc func() (ready bool, err error)

func wait(check CheckFunc) error {
	// load timeout values and such from config
	// call check func on a ticker
	// if error return error, if ready return nil
	timeout := time.After(10 * time.Minute)
	// configuration := config.Configuration()
	tick := time.Tick(1 * time.Second)
	ready := false

	for !ready {
		select {
		case <-timeout:
			return fmt.Errorf("there was a timeout")
		case <-tick:
			result, err := check()
			if err != nil {
				return err
			}
			ready = result
		}
	}
	return nil
}
