// Пишем барьер
package main

import (
	bar "concurrency/internal/trafficlight/barrier"
	"fmt"
	"sync"
	"time"
)

func main() {
	const nWorkers = 4
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(nWorkers)

	b := bar.NewBarrier(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func(i int) {
			// эмулируем подготовительный шаг...
			dur := time.Duration((i+1)*10) * time.Millisecond
			time.Sleep(dur)
			fmt.Printf("ready to go after %d ms\n", dur.Milliseconds())

			// ждем, пока все горутины соберутся у барьера
			b.Touch()

			// эмулируем основной шаг...
			fmt.Println("go!")
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("all done in %d ms\n", time.Since(start).Milliseconds())

	/*
		ready to go after 10 ms
		ready to go after 20 ms
		ready to go after 30 ms
		ready to go after 40 ms
		go!
		go!
		go!
		go!
		all done in 40 ms
	*/

}
