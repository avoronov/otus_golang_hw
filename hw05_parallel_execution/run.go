package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

// ErrErrorsLimitExceeded is the error which is thrown when appropriate amount of tasks failed.
var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// Task is the abstraction for piece of job to be done.
type Task func() error

func worker(taskCh <-chan Task, errorCh chan<- struct{}, quitCh <-chan struct{}) {
	var err error
	for {
		if err != nil {
			select {
			case errorCh <- struct{}{}:
				err = nil
			case <-quitCh:
				return
			}
		}

		select {
		case t := <-taskCh:
			err = t()
		case <-quitCh:
			return
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, workersCount int, maxErrors int) error {
	wg := sync.WaitGroup{}

	taskCh := make(chan Task)
	errorCh := make(chan struct{})
	quitCh := make(chan struct{})

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(taskCh, errorCh, quitCh)
		}()
	}

	var i int
	var totalErrors int
	for {
		if i == len(tasks) || totalErrors >= maxErrors {
			break
		}

		t := tasks[i]

		select {
		case <-errorCh:
			totalErrors++
		case taskCh <- t:
			i++
		}
	}

	close(quitCh)

	wg.Wait()

	close(taskCh)
	close(errorCh)

	if totalErrors >= maxErrors {
		return ErrErrorsLimitExceeded
	}

	return nil
}
