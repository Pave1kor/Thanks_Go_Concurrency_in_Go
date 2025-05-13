package stack

import "sync/atomic"

// начало решения

type Stack struct {
    top atomic.Pointer[Node]
}

type Node struct {
    val  int
    next *Node
}

// Push добавляет значение на вершину стека.
func (s *Stack) Push(val int) {
    newNode := &Node{val: val} // Создаем новый узел
    for {
        oldTop := s.top.Load()
        newNode.next = oldTop // Связываем новый узел со старым top
        if s.top.CompareAndSwap(oldTop, newNode) {
            break
        }
    }
}

// Pop удаляет и возвращает вершину стека.
// Если стек пуст, возвращает false.
func (s *Stack) Pop() (int, bool) {
    for {
        oldTop := s.top.Load()
        if oldTop == nil {
            return 0, false // Стек пуст
        }
        newTop := oldTop.next
        if s.top.CompareAndSwap(oldTop, newTop) {
            return oldTop.val, true
        }
        // Если CAS не удался, повторяем цикл
    }
}
