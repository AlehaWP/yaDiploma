package accrual

import (
	"context"
	"sync"
)

type WorkersPool struct {
	jobCh chan interface{}
	wg    sync.WaitGroup
	once  sync.Once
}

func (l *WorkersPool) Put(ctx context.Context, o interface{}) {
	select {
	case <-ctx.Done():
		l.once.Do(l.close)
		break
	default:
		l.jobCh <- o
	}
}

func (l *WorkersPool) close() {
	close(l.jobCh)
	l.wg.Wait()
}

func NewWorkersPool(ctx context.Context, f func(context.Context, interface{}), numOfWorkers int) *WorkersPool {

	l := &WorkersPool{
		jobCh: make(chan interface{}, 100),
	}

	l.wg.Add(numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {

		go func() {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			for job := range l.jobCh {
				f(ctx, job)
			}
			l.wg.Done()
		}()
	}
	return l
}
