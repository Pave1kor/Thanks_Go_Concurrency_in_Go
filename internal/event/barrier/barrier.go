// Барьер синхронизации.
package barrier

import (
	"sync"
)

// начало решения

// Barrier представляет барьер синхронизации.
type Barrier struct {
	remaining int
	cond      *sync.Cond
}

func NewBarrier(n int) *Barrier {
	return &Barrier{
		remaining: n,
		cond:      sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Barrier) Touch() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	b.remaining--
	if b.remaining == 0 {
		b.cond.Broadcast()
	}
	for b.remaining > 0 {
		b.cond.Wait()
	}
}

// конец решения
