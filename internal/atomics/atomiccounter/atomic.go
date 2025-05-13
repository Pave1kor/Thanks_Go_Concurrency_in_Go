// Атомарный счетчик
package atomic

import "sync/atomic"

// начало решения

// Total представляет атомарный счетчик.
type Total struct {
	counter atomic.Int32
}

// Increment увеличивает счетчик на 1.
func (t *Total) Increment() {
	t.counter.Add(1)
}

// Value возвращает значение счетчика.
func (t *Total) Value() int {
	return int(t.counter.Load())
}

// конец решения

