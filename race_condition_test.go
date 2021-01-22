package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 1

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println(x)
}
