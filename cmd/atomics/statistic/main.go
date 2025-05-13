package main
import (
	"fmt"
	"sync"
	statistic "concurrency/internal/atomics/statistic"
)

func main() {
	const nConc = 4
	var wg sync.WaitGroup
	wg.Add(nConc)
	// Вызываем внешнюю систему из нескольких горутин.
	ext := statistic.NewExternal()
	for range nConc {
		go func() {
			defer wg.Done()
			for range 10 {
				ext.Call()
			}
		}()
	}

	wg.Wait()

	// Количество вызовов и время последнего вызова.
	fmt.Println("Calls:", ext.NumCalls())
	fmt.Println("Last call:", ext.LastCall().Format("15:04:05"))
	// Calls: 40
	// Last call: 15:04:05
}