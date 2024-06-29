package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	mu := sync.Mutex{}
	ch := make(chan struct{}, n)
	cntErrors := 0

	for _, t := range tasks {
		mu.Lock()
		if cntErrors >= m && m >= 0 {
			mu.Unlock()
			return ErrErrorsLimitExceeded
		}
		mu.Unlock()

		wg.Add(1)
		ch <- struct{}{}

		go func(t Task) {
			defer func() {
				<-ch
				wg.Done()
			}()

			err := t()

			mu.Lock()
			if err != nil {
				cntErrors++
			}
			mu.Unlock()
		}(t)
	}

	return nil
}
