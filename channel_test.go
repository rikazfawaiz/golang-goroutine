package belajar_golang_goroutine

// channel variabel untuk mengirimkan data dari goroutine ke goroutine lain (synchron)

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // membuat channel
	defer close(channel)

	// channel <- "rikaz" // mengirim data ke channel
	// data := <- channel // mengirim data dari channel ke variabel
	// fmt.Println(<- channel) // langsung print data dari channel

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Rikaz Fawaiz Channel"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel  <- "Rikaz Fawaiz Channel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data:= <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//parameter in out
func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Rikaz Fawaiz Channel"
}

func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel daya tampung data di channel
func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	go func(){
		channel <- "Rikaz"
		channel <- "Fawaiz"
		channel <- "Test"
	}()

	go func(){
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")

	// fmt.Println(cap(channel)) // panjang buffer
	// fmt.Println(len(channel)) // jumlah data buffer
}