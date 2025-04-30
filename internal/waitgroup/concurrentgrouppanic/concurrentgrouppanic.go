package concurentgrouppanic

import (
	"sync"
)

// начало решения

// ConcGroup выполняет присылаемую работу в отдельных горутинах.
type ConcGroup struct {
	wg  sync.WaitGroup
	err chan any
}

// NewConcGroup создает новый экземпляр ConcGroup.
func NewConcGroup() *ConcGroup {
	return &ConcGroup{wg: sync.WaitGroup{}, err: make(chan any, 1)}
}

// Run выполняет присланную работу в отдельной горутине.
// Если горутина запаниковала, Run не паникует.
func (p *ConcGroup) Run(work func()) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				select {
				case p.err <- r:
				default:
				}
			}
			p.wg.Done()
		}()
		work()
	}()
}

// Wait ожидает, пока не закончится вся выполняемая в данный момент работа.
// Если запаниковала хотя бы одна из горутин, запущенных через Run -
// Wait тоже паникует.
func (p *ConcGroup) Wait() {
	p.wg.Wait()
	select {
	case val := <-p.err:
		p.err <- val
		panic(val)
	default:
	}
}

// конец решения
