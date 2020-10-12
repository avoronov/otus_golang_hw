package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

// ErrErrorsLimitExceeded is the error which is thrown when appropriate amount of tasks failed.
var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// Task is the abstraction for piece of job to be done.
type Task func() error

func worker(taskach <-chan Task, errorach chan<- struct{}) {
	for {
		select {
		case t, ok := <-taskach:
			if !ok {
				return
			}
			err := t()
			if err != nil {
				errorach <- struct{}{}
			}
			_ = 1 // DELETE ME
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	var totalErrors, totalWorkers int
	var wg sync.WaitGroup

	taskach := make(chan Task)
	errorach := make(chan struct{})

	for i, t := range tasks {
		if totalErrors >= M {
			break
		}

		if (i-totalErrors >= totalWorkers) && totalWorkers < N {
			go func() {
				wg.Add(1)
				defer wg.Done()
				worker(taskach, errorach)
			}()
			totalWorkers++
		}

		for {
			var stop bool

			select {
			case <-errorach:
				totalErrors++
				if totalErrors >= M {
					stop = true
				}
			case taskach <- t:
				stop = true
			default:
			}

			if stop {
				break
			}
		}
	}

	close(taskach)

	wg.Wait()

	close(errorach)

	if totalErrors >= M {
		return ErrErrorsLimitExceeded
	}

	return nil
}
