package main

import (
	runtime "concurrency/internal/runtimesimulator"
	"fmt"
)

func main() {
	// создаем рантайм на 2 потока
	r := runtime.NewRuntime(2)

	// создаем 4 горутины
	g1 := r.Go()
	g2 := r.Go()
	r.Go()
	r.Go()
	r.Schedule()

	// прошло 10 единиц времени, g1 завершила выполнение, g2 заблокирована
	r.Forward(10)
	g1.Done()
	g2.Block()
	r.Schedule()

	// выводим текущее состояние рантайма
	state := r.State()
	fmt.Printf("%+v\n", state)
	// {dur:10 threads:map[1:3 2:4] runnable:[] running:[3 4] waiting:[2] dead:[1]}
}
