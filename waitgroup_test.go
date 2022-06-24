package belajar_golang_goroutine

//wait group menunggu selesainya goroutine running ***

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func TestWaitingGroup (t *testing.T) {
	group := &sync.WaitGroup{}

	for i:=0; i<100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Completed")
}