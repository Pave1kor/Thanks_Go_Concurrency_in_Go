package timeit

import "time"

func Timeit(nIter int, nWorkers int, fn func()) int {
	done := make(chan struct{}, nWorkers)
	start := time.Now()

	// работают nWorkers параллельных обработчиков
	for i := 0; i < nWorkers; i++ {
		go func() {
			// каждый обработчик выполняет nIter/nWorkers итераций
			for i := 0; i < nIter/nWorkers; i++ {
				fn()
			}
			done <- struct{}{}
		}()
	}

	// дожидаемся завершения обработчиков
	for i := 0; i < nWorkers; i++ {
		<-done
	}

	return int(time.Since(start).Milliseconds())
}

// конец решения
