package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0

	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println(x)
}
