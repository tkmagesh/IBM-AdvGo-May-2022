package runner

import (
	"errors"
	"time"
)

type Task func(int)

var ErrTimeout = errors.New("timeout occured")

type Runner struct {
	/*  */
}

func (r *Runner) Add(task Task) {
	/*  */
}

func (r *Runner) Start() error {
	/*  */
	return nil
}

func New(timeout time.Duration) *Runner {
	return &Runner{
		/*  */
	}
}
