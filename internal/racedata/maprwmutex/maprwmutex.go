package maprwmutex

import "sync"

// начало решения

// Counter представляет безопасную карту частот слов.
// Ключ - строка, значение - целое число.
type Counter struct {
	// не меняйте название и тип поля lock
	lock  sync.RWMutex
	count map[string]int
	// ...
}

// Increment увеличивает значение по ключу на 1.
func (c *Counter) Increment(str string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count[str]++
	// ...
}

// Value возвращает значение по ключу.
func (c *Counter) Value(str string) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.count[str]
	// ...
}

// Range проходит по всем записям карты,
// и для каждой вызывает функцию fn, передавая в нее ключ и значение.
func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	for key, val := range c.count {
		fn(key, val)
	}
	// ...
}

// NewCounter создает новую карту частот.
func NewCounter() *Counter {
	return &Counter{
		count: make(map[string]int),
		lock:  sync.RWMutex{},
	}
	// ...
}

// конец решения
