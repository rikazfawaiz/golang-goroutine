package belajar_golang_goroutine

// race condition : dimana satu variabel digunakan lebih dari 1 goroutine menyebabkan data dari beberapa goroutine nilainya sama

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i:=1; i<=1000; i++ {
		go func(){
			for j:=1; i<=100; j++ {
				x++
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ",x)
}