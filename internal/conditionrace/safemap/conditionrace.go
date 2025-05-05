package conditionrace

import (
	"sync"
)

// ConcMap - безопасная в многозадачной среде карта.
type ConcMap[K comparable, V any] struct {
	items map[K]V
	lock  sync.Mutex
}

// NewConcMap создает новую карту.
func NewConcMap[K comparable, V any]() *ConcMap[K, V] {
	return &ConcMap[K, V]{items: map[K]V{}}
}

// Get возвращает значение по ключу.
func (cm *ConcMap[K, V]) Get(key K) V {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	return cm.items[key]
}

// Set устанавливает значение по ключу.
func (cm *ConcMap[K, V]) Set(key K, val V) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.items[key] = val
}

// SetIfAbsent устанавливает новое значение по ключу
// и возвращает его, но только если такого ключа нет в карте.
// Если ключ уже есть - возвращает старое значение по ключу.
func (cm *ConcMap[K, V]) SetIfAbsent(key K, val V) V {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if _, ok := cm.items[key]; !ok {
		cm.items[key] = val
		return val
	}
	return cm.items[key]
}

// Compute устанавливает значение по ключу, применяя к нему функцию.
// Возвращает новое значение. Функция выполняется атомарно.
func (cm *ConcMap[K, V]) Compute(key K, f func(V) V) V {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	v := cm.items[key]
	v = f(v)
	cm.items[key] = v
	return v
	// TODO: реализовать требования
}
