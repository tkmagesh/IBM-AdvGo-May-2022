package main

import (
	"fmt"
	"math/rand"
	"time"
	"worker-demo/worker"
)

type MyWork struct {
	id int
}

/* implementation of the worker.Work interface */
func (m MyWork) Task() {
	fmt.Println("task started - ", m.id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println("task completed - ", m.id)
}

func main() {
	w := worker.New(5)
	for i := 1; i <= 20; i++ {
		w.Run(MyWork{id: i})
	}
	fmt.Println("All tasks are assigned")
	w.Shutdown()
	//w.Run(MyWork{id: 100}) //SHOULD NOT BE ALLOWED as the worker is already "Shutdown"
}
