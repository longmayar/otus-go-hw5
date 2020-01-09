package hw5

import (
	"errors"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Run(tasks []func() error, routinesLimit int, errLimit int) error {
	in := make(chan func() error, len(tasks))
	out := make(chan error, len(tasks))
	quit := make(chan interface{})

	for i := 0; i < routinesLimit; i++ {
		wg.Add(1)
		go worker(in, out, quit)
	}

	go limiter(out, quit, errLimit)

	for _, task := range tasks {
		in <- task
	}

	wg.Wait()
	close(out)

	select {
	case <-quit:
		fmt.Println("ret error")
		return errors.New("error limit is exceeded")
	default:
		fmt.Println("ret nil")
		return nil
	}
}

func worker(in chan func() error, out chan error, quit chan interface{}) {
	defer func() {
		fmt.Println("wg done")
		wg.Done()
	}()

	for {
		select {
		case <-quit:
			fmt.Println("quit1")
			return
		default:
		}

		select {
		case task := <-in:
			fmt.Println("from in to out")
			out <- task()
		case <-quit:
			fmt.Println("quit2")
			return
		}
	}
}

func limiter(out chan error, quit chan interface{}, errLimit int) {
	var errCnt int

	for res := range out {
		fmt.Println("from out")

		if res != nil {
			errCnt++
			fmt.Println("err", errCnt)
		}

		if errCnt == errLimit {
			fmt.Println("close")
			close(quit)
			return
		}
	}
}
