package wererace1

import "time"

func Delay(duration time.Duration, fn func()) func() {
	canceled := false // (1)

	go func() {
		time.Sleep(duration)
		if !canceled { // (2)
			fn()
		}
	}()

	cancel := func() {
		canceled = true // (3)
	}
	return cancel // (4)
}
