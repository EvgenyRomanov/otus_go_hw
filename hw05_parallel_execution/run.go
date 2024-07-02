package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ch := make(chan struct{}, n)
	var cntErrors int64

	for _, t := range tasks {
		if atomic.LoadInt64(&cntErrors) >= int64(m) && m >= 0 {
			return ErrErrorsLimitExceeded
		}

		wg.Add(1)
		ch <- struct{}{}

		go func(t Task) {
			defer func() {
				<-ch
				wg.Done()
			}()

			err := t()
			if err != nil {
				atomic.AddInt64(&cntErrors, 1)
			}
		}(t)
	}

	return nil
}
