package limiter

import (
	"errors"
	"sync"
	"time"
)

var ErrBusy = errors.New("busy")
var ErrCanceled = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func Throttle(limit int, fn func()) (handle func() error, cancel func()) {
	ch := make(chan struct{}, limit)
	var err error
	var mu sync.Mutex
	open := true
	ticker := time.NewTicker(time.Second)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	handle = func() error {
		select {
		case <-ticker.C:
			mu.Lock()
			open = true
			mu.Unlock()
		case _, ok := <-ch:
			go func() {
				if ok && !open {
					fn()
					ch <- struct{}{}
					err = nil
				} else {
					err = ErrCanceled
				}
			}()
			return err
		default:
			err = ErrBusy
			open = false
		}
		return err
	}

	cancel = func() {
		ticker.Stop()
		close(ch)
	}
	return
}

// конец решения
