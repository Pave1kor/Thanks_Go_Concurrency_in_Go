// Статистика вызовов.
package statistic

import (
	"sync/atomic"
	"time"
)

// начало решения

// External представляет внешний сервис.
type External struct {
	lastCall atomic.Value
	numCalls atomic.Int32
}

// NewExternal создает новый экземпляр External.
func NewExternal() *External {
	return &External{}
}

// Call вызывает внешний сервис.
func (e *External) Call() {
	// вызываем внешний сервис...
	e.lastCall.Store(time.Now())
	e.numCalls.Add(1)
}

// LastCall возвращает время последнего вызова.
func (e *External) LastCall() time.Time {
	return e.lastCall.Load().(time.Time)
}

// NumCalls возвращает количество вызовов.
func (e *External) NumCalls() int {
	return int(e.numCalls.Load())
}

// конец решения
