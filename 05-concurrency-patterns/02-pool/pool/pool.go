package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrInvalidPoolSize = errors.New("invalid pool size. pool size has to be > 0")
var ErrPoolClosed = errors.New("cannot aquire resources when the pool is closed")

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	mutex     *sync.Mutex
}

func (p *Pool) Acquire() (io.Closer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire: From Pool")
		return r, nil
	default:
		fmt.Println("Acquire: From Factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case p.resources <- resource:
		fmt.Println("Release : Into the pool")
		return nil
	default:
		fmt.Println("Release: Close & discard the resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}

func New(factory func() (io.Closer, error), poolSize int) (*Pool, error) {
	if poolSize <= 0 {
		return nil, ErrInvalidPoolSize
	}
	pool := &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		closed:    false,
		mutex:     &sync.Mutex{},
	}

	return pool, nil
}
