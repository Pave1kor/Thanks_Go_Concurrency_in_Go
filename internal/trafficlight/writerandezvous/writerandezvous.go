// Пишем рандеву
package writerandezvous

// начало решения

// Rendezvous представляет рандеву двух горутин.
type Rendezvous struct {
	ch chan struct{}
}

// NewRendezvous создает новое рандеву.
func NewRendezvous() *Rendezvous {
	ch := make(chan struct{}, 2)
	ch <- struct{}{}
	ch <- struct{}{}
	return &Rendezvous{
		ch: ch,
	}
}

// Ready фиксирует, что вызывающая горутина прибыла к точке сбора.
// Блокирует вызывающую горутину, пока не прибудет вторая.
// Когда обе горутины прибудут, Ready их разблокирует.
func (r *Rendezvous) Ready() {
	<-r.ch
	for len(r.ch) != 0 {
	}
}

// // Rendezvous представляет рандеву двух горутин.
// type Rendezvous struct {
// 	wg *sync.WaitGroup
// }

// // NewRendezvous создает новое рандеву.
// func NewRendezvous() *Rendezvous {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	return &Rendezvous{&wg}
// }

// // Ready фиксирует, что вызывающая горутина прибыла к точке сбора.
// // Блокирует вызывающую горутину, пока не прибудет вторая.
// // Когда обе горутины прибудут, Ready их разблокирует.
// func (r *Rendezvous) Ready() {
// 	r.wg.Done()
// 	r.wg.Wait()
// }
