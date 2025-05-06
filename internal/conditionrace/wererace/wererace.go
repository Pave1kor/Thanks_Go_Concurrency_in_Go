package wererace

import "time"

func Delay(duration time.Duration, fn func()) func() {
	canceled := make(chan struct{}) // (1)

	go func() {
		timer := time.NewTimer(duration)
		select {
		case <-timer.C:
			fn() // (2)
		case <-canceled:
			timer.Stop() // (3)
		}
	}()

	return func() {
		select {
		case <-canceled:
		default:
			close(canceled) // (4)
		}
	}
}
