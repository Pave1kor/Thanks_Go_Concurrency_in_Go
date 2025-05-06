package limiter

import (
	"errors"
	"sync"
)

var ErrBusy = errors.New("busy")
var ErrCanceled = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func Throttle(limit int, fn func()) (handle func() error, cancel func()) {
    tokens := make(chan struct{}, limit)
    for range limit {
        tokens <- struct{}{}
    }

    var (
        mu     sync.Mutex
        closed bool
    )

    handle = func() error {
        mu.Lock()
        if closed {
            mu.Unlock()
            return ErrCanceled
        }
        mu.Unlock()

        select {
        case <-tokens:
            go func() {
                defer func() { tokens <- struct{}{} }()
                fn()
            }()
            return nil
        default:
            return ErrBusy
        }
    }

    cancel = func() {
        mu.Lock()
        closed = true
        mu.Unlock()
    }

    return
}

// конец решения
