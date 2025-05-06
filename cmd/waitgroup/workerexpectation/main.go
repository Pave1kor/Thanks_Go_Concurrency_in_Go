package main

import (
	work "concurrency/internal/waitgroup/workerexpectation"
	"errors"
	"fmt"
	"time"
)

func main() {
	{
		// Завершение по ошибке
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			if count == 0 {
				return errors.New("count is zero")
			}
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := work.NewWorker(fn)
		worker.Start()
		time.Sleep(25 * time.Millisecond)

		fmt.Println()
		// 3 2 1
	}
	{
		// Завершение по Stop
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := work.NewWorker(fn)
		worker.Start()
		time.Sleep(25 * time.Millisecond)
		worker.Stop()

		fmt.Println()
		// 3 2 1
	}
	{
		// Ожидание завершения через Wait
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := work.NewWorker(fn)
		worker.Start()

		// эта горутина остановит работягу через 25 мс
		go func() {
			time.Sleep(25 * time.Millisecond)
			worker.Stop()
		}()

		// подождем, пока кто-нибудь остановит работягу
		worker.Wait()
		fmt.Println("done")

		// 3 2 1 done
	}
}
