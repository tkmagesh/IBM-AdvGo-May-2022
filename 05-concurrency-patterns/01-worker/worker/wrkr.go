package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

/* Implement the following methods */

type Worker struct {
	/* fill in the blanks */
	workQueue chan Work
	wg        sync.WaitGroup
}

func (w *Worker) Run(work Work) {
	/* fill in the blanks */
	w.workQueue <- work
}

func (w *Worker) Shutdown() {
	/* fill in the blanks */
	close(w.workQueue)
	w.wg.Wait()
	fmt.Println("Worker shutdown completed")
}

func New(workerCount int) *Worker {
	worker := &Worker{
		workQueue: make(chan Work),
	}
	worker.wg.Add(workerCount)
	for idx := 0; idx < workerCount; idx++ {
		go func(id int) {
			for w := range worker.workQueue {
				fmt.Println("Worker Id : ", id)
				w.Task()
			}
			worker.wg.Done()
		}(idx)
	}
	return worker

}
