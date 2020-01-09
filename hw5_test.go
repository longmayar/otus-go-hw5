package hw5

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunChan(t *testing.T) {
	successTasks := 0
	errorTasks := 0
	resChan := make(chan error, 9)

	task1 := func() error { err := errors.New("error task 1"); resChan <- err; return err }
	task2 := func() error { err := errors.New("error task 2"); resChan <- err; return err }
	task3 := func() error { err := errors.New("error task 3"); resChan <- err; return err }
	task4 := func() error { err := errors.New("error task 4"); resChan <- err; return err }
	task5 := func() error { err := errors.New("error task 5"); resChan <- err; return err }
	task6 := func() error { err := errors.New("error task 6"); resChan <- err; return err }
	task7 := func() error { err := errors.New("error task 7"); resChan <- err; return err }
	task8 := func() error { err := errors.New("error task 8"); resChan <- err; return err }
	task9 := func() error { err := errors.New("error task 9"); resChan <- err; return err }

	tasks := []func() error{task1, task2, task3, task4, task5, task6, task7, task8, task9}
	result := Run(tasks, 2, 2)

	for i := 0; i < len(tasks); i++ {
		resTask := <-resChan
		if resTask != nil {
			errorTasks++
		} else {
			successTasks++
		}
	}

	assert.LessOrEqual(t, successTasks, 4)
	assert.Equal(t, 2, errorTasks)
	expected := errors.New("error limit is exceeded")
	assert.Equal(t, expected, result)
}
