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
	tasksCh := make(chan Task, n)
	defer close(tasksCh)

	var cntErrors int64

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range tasksCh {
				err := task()
				if atomic.LoadInt64(&cntErrors) >= int64(m) && m >= 0 {
					return
				}
				if err != nil {
					atomic.AddInt64(&cntErrors, 1)
				}
			}
		}()
	}

	for _, t := range tasks {
		tasksCh <- t
		if atomic.LoadInt64(&cntErrors) >= int64(m) && m >= 0 {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
