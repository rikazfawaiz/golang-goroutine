package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"sync"
	"sync/atomic"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i:=1; i<=1000; i++ {
		go func(){
			group.Add(1)
			for j:=1; i<=100; j++ {
				atomic.AddInt64(&x,1) // untuk manipulasi tipe data primitif, gak perlu pake mutex
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter = ",x)
}