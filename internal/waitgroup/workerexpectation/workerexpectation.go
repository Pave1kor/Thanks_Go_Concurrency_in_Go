package workerexpectation

import (
	"sync"
)

// Worker выполняет заданную функцию в цикле, пока не будет остановлен.
type Worker struct {
	fn      func() error
	wg      sync.WaitGroup
	mu      sync.Mutex
	running bool
	stopped bool
}

// NewWorker создает новый экземпляр Worker с заданной функцией.
func NewWorker(fn func() error) *Worker {
	return &Worker{fn: fn}
}

// Start запускает отдельную горутину, в которой циклически
// выполняет заданную функцию, пока не будет вызван метод Stop,
// либо пока функция не вернет ошибку.
// Повторные вызовы Start игнорируются.
// Гарантируется, что Start не вызывается из разных горутин.
func (w *Worker) Start() {
	w.mu.Lock()
	if w.running {
		w.mu.Unlock()
		return
	}
	w.running = true
	w.stopped = false
	w.mu.Unlock()

	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			w.mu.Lock()
			stopped := w.stopped
			w.mu.Unlock()
			if stopped {
				break
			}

			err := w.fn()
			if err != nil {
				break
			}
		}
	}()
}

// Stop останавливает выполнение цикла.
// Вызов Stop до Start игнорируется.
// Повторные вызовы Stop игнорируются.
// Гарантируется, что Stop не вызывается из разных горутин.
func (w *Worker) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if !w.running || w.stopped {
		return
	}
	w.stopped = true
}

// Wait блокирует вызвавшую его горутину до тех пор,
// пока Worker не будет остановлен (из-за ошибки или вызова Stop).
// Wait может вызываться несколько раз, в том числе из разных горутин.
// Wait может вызываться до Start. Это не приводит к блокировке.
// Wait может вызываться после Stop. Это не приводит к блокировке.
func (w *Worker) Wait() {
	w.wg.Wait()
}
