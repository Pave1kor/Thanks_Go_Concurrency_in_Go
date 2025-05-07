package tryacquire

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
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// TryAcquire занимает место в семафоре, если есть свободное,
// и возвращает true. В противном случае просто возвращает false.
func (s Semaphore) TryAcquire() bool {
	if len(s.ch) < cap(s.ch) {
		s.Acquire()
		return true
	}
	return false
}

// Release освобождает место в семафоре и разблокирует
// одну из заблокированных горутин (если такие были).
func (s *Semaphore) Release() {
	<-s.ch
}
