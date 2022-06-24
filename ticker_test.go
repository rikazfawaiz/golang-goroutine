package belajar_golang_goroutine

//kejadian berulang

import (
	"fmt"
	"testing"
	"time"
	// "sync"
)

func TestTicker(t *testing.T){
	ticker := time.NewTicker(1 * time.Second)

	go func(){
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time :=  range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T){ //hanya mengembalikan channelnya saja
	channel := time.Tick(1 * time.Second)

	for time :=  range channel {
		fmt.Println(time)
	}
}