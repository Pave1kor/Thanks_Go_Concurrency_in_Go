package timeit

import (
	"sync"
	"time"
)

func Timeit(nIter int, nWorkers int, fn func()) int {
	var wg sync.WaitGroup
	start := time.Now()

	// работают nWorkers параллельных обработчиков
	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			// каждый обработчик выполняет nIter/nWorkers итераций
			for i := 0; i < nIter/nWorkers; i++ {
				fn()
			}
			wg.Done()
		}()
	}

	// дожидаемся завершения обработчиков
	wg.Wait()
	return int(time.Since(start).Milliseconds())
}

// конец решения
