package wait

type CheckFunc func() (ready bool, err error)

func wait(check CheckFunc) error {
	// load timeout values and such from config
	// call check func on a ticker
	// if error return error, if ready return nil
	return nil
}
