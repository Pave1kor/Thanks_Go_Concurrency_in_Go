package query

import "sync"

// начало решения

// Queue - блокирующая FIFO-очередь.
type Queue struct {
	// TODO: переделать на срез и sync.Cond.
	cond  sync.Cond
	items []int
}

// NewQueue создает новую очередь.
func NewQueue() *Queue {
	// TODO: очередь должна быть безразмерной.
	return &Queue{
		items: make([]int, 0),
		cond:  *sync.NewCond(&sync.Mutex{}),
	}
}

// Put добавляет элемент в очередь.
// Поскольку очередь безразмерная, никогда не блокируется.
// This function adds an item to the queue
func (q *Queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.items = append(q.items, item)
	q.cond.Signal()

}

// Get извлекает элемент из очереди.
// Если очередь пуста, блокируется до момента,
// пока в очереди не появится элемент.
func (q *Queue) Get() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Len возвращает количество элементов в очереди.
func (q *Queue) Len() int {
	return len(q.items)
}

// конец решения
