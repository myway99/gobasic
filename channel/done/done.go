package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker)  {
	for n := range w.in {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c\n",
			id, n)
		// go func() {done <- true}()
		w.done()
	}
}

type worker struct {
	in	chan	int
	//wg	*sync.WaitGroup
	done func()
}

func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo()  {
	//var c chan int  // c == nil
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		// <-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		// <-workers[i].done
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
