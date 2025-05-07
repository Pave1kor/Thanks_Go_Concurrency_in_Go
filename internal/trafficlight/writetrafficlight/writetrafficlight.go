package writetrafficlight

// начало решения

// Semaphore представляет семафор синхронизации.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore создает новый семафор указанной вместимости.
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, n)}
}

// Acquire занимает место в семафоре, если есть свободное.
// В противном случае блокирует вызывающую горутину.
// This function is used to acquire a semaphore
func (s *Semaphore) Acquire() {
	// Send an empty struct to the channel
	s.ch <- struct{}{}
}

// Release освобождает место в семафоре и разблокирует
// одну из заблокированных горутин (если такие были).
func (s *Semaphore) Release() {
	<-s.ch
}

// конец решения
