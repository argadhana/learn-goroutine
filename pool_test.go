package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Argadhana")
	pool.Put("Arga")
	pool.Put("Nugroho")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	fmt.Println("SELESEI")
	group := &sync.WaitGroup{}

	group.Wait()
}
