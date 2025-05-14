// Симулятор рантайма Go.
package runtimesimulator

// максимальное время непрерывного выполнения горутины на потоке
const maxRunDur = 100

// статус горутины
type gStatus string

// статусы горутин
var (
	statusRunnable gStatus = "runnable" // готова к выполнению
	statusRunning  gStatus = "running"  // выполняется на потоке
	statusWaiting  gStatus = "waiting"  // заблокирована
	statusDead     gStatus = "dead"     // завершилась
)

// Goroutine представляет горутину.
type Goroutine struct {
	id     int     // идентификатор, нумерация с 1
	runDur int     // время в состоянии running
	status gStatus // статус
}

// Block переводит горутину в состояние waiting.
func (g *Goroutine) Block() {
	if g.status != statusRunning {
		panic("invalid status for block: " + g.status)
	}
	g.status = statusWaiting
}

// Unblock переводит горутину в состояние runnable.
func (g *Goroutine) Unblock() {
	if g.status != statusWaiting {
		panic("invalid status for unblock: " + g.status)
	}
	g.status = statusRunnable
}

// Done переводит горутину в состояние dead.
func (g *Goroutine) Done() {
	if g.status != statusRunning {
		panic("invalid status for done: " + g.status)
	}
	g.status = statusDead
}

// Thread представляет поток операционной системы.
type Thread struct {
	id   int        // идентификатор, нумерация с 1
	goro *Goroutine // горутина на выполнении
}

// RuntimeState представляет состояние рантайма.
type RuntimeState struct {
	dur      int         // общее время выполнения
	threads  map[int]int // ключ - id потока, значение - id горутины (0 - поток свободен)
	runnable []int       // id горутин в очереди на выполнение
	running  []int       // id горутин на выполнении
	waiting  []int       // id заблокированных горутин
	dead     []int       // id завершенных горутин
}

// Runtime представляет симулятор рантайма.
type Runtime struct {
	dur      int          // общее время выполнения
	nGoro    int          // счетчик горутин
	threads  []*Thread    // потоки
	runnable []*Goroutine // горутины в очереди на выполнение
	running  []*Goroutine // горутины на выполнении
	waiting  []*Goroutine // заблокированные горутины
	dead     []*Goroutine // завершенные горутины
}

// NewRuntime создает новый рантайм на gomaxprocs потоках.
func NewRuntime(gomaxprocs int) *Runtime {
	threads := make([]*Thread, gomaxprocs)
	for i := range gomaxprocs {
		threads[i] = &Thread{id: i + 1}
	}
	return &Runtime{threads: threads}
}

// Go создает новую горутину в рантайме.
func (r *Runtime) Go() *Goroutine {
	r.nGoro++
	g := &Goroutine{id: r.nGoro, status: statusRunnable}
	r.runnable = append(r.runnable, g)
	return g
}

// начало решения

// addGoroutinesToThreads добавляет горутины из runnable в свободные потоки
func (r *Runtime) addGoroutinesToThreads() {
	for _, t := range r.threads {
		if t.goro == nil && len(r.runnable) > 0 {
			// Берем первую горутину из очереди
			g := r.runnable[0]
			r.runnable = r.runnable[1:]

			// Назначаем горутину на поток
			t.goro = g
			g.status = statusRunning
			g.runDur = 0

			// Добавляем в running, если еще не там
			found := false
			for _, rg := range r.running {
				if rg.id == g.id {
					found = true
					break
				}
			}
			if !found {
				r.running = append(r.running, g)
			}
		}
	}
}

// Forward двигает время вперед на dur единиц.
func (r *Runtime) Forward(dur int) {
	r.dur += dur

	// Добавляем горутины в свободные потоки
	r.addGoroutinesToThreads()

	// Увеличиваем время выполнения горутин в состоянии running
	for _, g := range r.running {
		g.runDur += dur
	}
}

// Schedule планирует выполнение горутин.
func (r *Runtime) Schedule() {
	// Сначала обрабатываем разблокированные горутины в waiting
	newWaiting := make([]*Goroutine, 0, len(r.waiting))
	for _, g := range r.waiting {
		if g.status == statusRunnable {
			r.runnable = append(r.runnable, g) // Переносим в runnable
		} else {
			newWaiting = append(newWaiting, g) // Оставляем в waiting
		}
	}
	r.waiting = newWaiting

	// Затем обрабатываем потоки
	for _, t := range r.threads {
		if t.goro == nil {
			continue
		}

		g := t.goro

		switch {
		case g.runDur >= maxRunDur:
			g.status = statusRunnable
			r.runnable = append(r.runnable, g)
			t.goro = nil

		case g.status == statusWaiting:
			// Проверяем, не была ли горутина уже разблокирована
			r.waiting = append(r.waiting, g)
			t.goro = nil

		case g.status == statusDead:
			r.dead = append(r.dead, g)
			t.goro = nil

		case g.status == statusRunnable:
			r.runnable = append(r.runnable, g)
		}
	}

	// Обновляем список running
	r.running = r.running[:0]
	for _, t := range r.threads {
		if t.goro != nil {
			r.running = append(r.running, t.goro)
		}
	}

	// Добавляем горутины в свободные потоки
	r.addGoroutinesToThreads()
}

// конец решения

// State возвращает текущее состояние рантайма.
func (r *Runtime) State() RuntimeState {
	threads := make(map[int]int)
	for _, t := range r.threads {
		if t.goro != nil {
			threads[t.id] = t.goro.id
		} else {
			threads[t.id] = 0
		}
	}
	runnable := make([]int, len(r.runnable))
	for i, g := range r.runnable {
		runnable[i] = g.id
	}
	running := make([]int, len(r.running))
	for i, g := range r.running {
		running[i] = g.id
	}
	waiting := make([]int, len(r.waiting))
	for i, g := range r.waiting {
		waiting[i] = g.id
	}
	dead := make([]int, len(r.dead))
	for i, g := range r.dead {
		dead[i] = g.id
	}
	return RuntimeState{
		dur:      r.dur,
		threads:  threads,
		runnable: runnable,
		running:  running,
		waiting:  waiting,
		dead:     dead,
	}
}
