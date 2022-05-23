package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {

	var group = sync.WaitGroup{}
	var x int64 = 0
	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x , 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("Perulangen ke = ", x)

}