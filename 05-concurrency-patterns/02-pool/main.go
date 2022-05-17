package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pool-demo/pool"
	"sync"
	"time"
)

//Resource
type DBConnection struct {
	ID int
}

func (dbConnection *DBConnection) Close() error {
	fmt.Printf("Closing and discarding the resource (ID) : %d\n", dbConnection.ID)
	return nil
}

//Factory
var idCounter int

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	fmt.Printf("DBConnectionFactory : Creating resource (ID) : %d\n", idCounter)
	return &DBConnection{ID: idCounter}, nil
}

func main() {
	/*
		create an instance of a Pool (with the pool size and a factory function)

		What is a resource?
			Any object that implements io.closer interface ( close() method )

		When a resource is Acquired?
			the pool will check if it has any resources in the resource-pool
			if yes, return the resource from the resource-pool
			else create a new resource using the factory and return it

		When a resouce is Released?
			the pool will check if it is full
			if yes, the resource will be discarded (by calling the Close() method)
			else maintain the released resource in the resource-pool

		when the pool is "Close()"?
			prevent anymore acquisition of the resources
			make sure all the resources are 'closed' and discard them

		Important Note:
			When a resource is "Acquired" by a client, the same resource SHOULD NOT BE given to another client

		APIs
			-New(factory, poolSize)
			pool
				-Acquire() => resource
				-Release(resource)
				-Close()

	*/

	p, err := pool.New(DBConnectionFactory, 5)

	if err != nil {
		log.Fatalln(err)
	}

	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 0; client < clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()

	fmt.Println("Second batch of operations. Hit ENTER to start")
	var input string
	fmt.Scanln(&input)
	wg.Add(3)
	for client := 0; client < 3; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Worker : %d, Acquired: %d\n", id, conn.(*DBConnection).ID)
	time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
	fmt.Printf("Worker %d Done. Releasing resource : %d\n", id, conn.(*DBConnection).ID)
	p.Release(conn)
}
