package wscore

import (
	"sync"
)

var MyWorker *Pool

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(max int) *Pool {
	p := &Pool{work: make(chan Worker)}

	p.wg.Add(max)
	for i := 0; i < max; i++ {
		go func() {
			for worker := range p.work {
				worker.Task()
			}
			p.wg.Done()
		}()
	}

	return p
}

func (p *Pool) Run(work Worker) {
	p.work <- work
}
