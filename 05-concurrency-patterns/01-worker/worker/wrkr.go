package worker

type Work interface {
	Task()
}

/* Implement the following methods */

type Worker struct {
	/* fill in the blanks */
}

func (w *Worker) Run(work Work) {
	/* fill in the blanks */
	work.Task()
}

func (w *Worker) Shutdown() {
	/* fill in the blanks */
}

func New(workerCount int) *Worker {
	return &Worker{}
}
