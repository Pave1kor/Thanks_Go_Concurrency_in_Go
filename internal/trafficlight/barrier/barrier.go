package barrier

import "sync"

// Barrier представляет барьер синхронизации.
type Barrier struct {
	wg *sync.WaitGroup
}

// NewBarrier создает новый барьер с указанным порогом.
func NewBarrier(n int) *Barrier {
	wg := sync.WaitGroup{}
	wg.Add(n)
	return &Barrier{
		wg: &wg,
	}
}

// Touch фиксирует, что вызывающая горутина достигла барьера.
// Если барьера достигли меньше n горутин, Touch блокирует вызывающую горутину.
// Когда n горутин достигнут барьера, Touch разблокирует их все.
func (b *Barrier) Touch() {
	b.wg.Done()
	b.wg.Wait()
}
