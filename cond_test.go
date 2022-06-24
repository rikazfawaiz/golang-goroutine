package belajar_golang_goroutine

// cond, condition dimana ginukan jika ada kondisi tertentu setelah melakukan Lock

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i:=0; i<10; i++ {
		go WaitCondition(i)
	}

	go func(){ // satu" jalan setelah signal
		for i:=0; i<10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}	
	}()

	// go func(){ // jalankan semua
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}