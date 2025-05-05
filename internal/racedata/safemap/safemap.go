package safemap

import "sync"

// начало решения

// Counter представляет безопасную карту частот слов.
// Ключ - строка, значение - целое число.
type Counter struct {
	count map[string]int
	mu    sync.Mutex
}

// Increment увеличивает значение по ключу на 1.
func (c *Counter) Increment(str string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count[str]++
}

// Value возвращает значение по ключу.
func (c *Counter) Value(str string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count[str]
}

// Range проходит по всем записям карты,
// и для каждой вызывает функцию fn, передавая в нее ключ и значение.
func (c *Counter) Range(fn func(key string, val int)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.count {
		fn(key, val)
	}
}

// NewCounter создает новую карту частот.
func NewCounter() *Counter {
	return &Counter{
		count: make(map[string]int),
		mu:    sync.Mutex{},
	}
}
