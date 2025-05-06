package wererace2

import "time"

func Delay(duration time.Duration, fn func()) func() {
	alive := make(chan struct{}) // (1)
	close(alive)                 // (2)

	go func() {
		time.Sleep(duration)
		select {
		case <-alive: // (3)
			fn()
		default:
		}
	}()

	cancel := func() {
		alive = nil // (4)
	}
	return cancel
}
