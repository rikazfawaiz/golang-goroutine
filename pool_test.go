package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New : func() interface{} { //bikin default value biar gak nil
			return "New"
		},
	}
	// group := sync.WaitGroup{}

	pool.Put("Rikaz")
	pool.Put("Fawaiz")
	pool.Put("Haerul")
	pool.Put("Afgani")

	for i:=0; i<10; i++ {
		go func(){
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	// group.Wait()
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}