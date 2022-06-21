package belajar_golang_goroutine

// channel variabel untuk mengirimkan data dari goroutine ke goroutine lain (synchron)

import (
	"fmt"
	"testing"
	"time"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func(){
		for i:=0; i<10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data",data)
	}

	fmt.Println("Selesai")
}

//select data lebih dari 1 channel
func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <- channel1:
				fmt.Println("Channel 1 ",data)
				counter++
			case data := <- channel2:
				fmt.Println("Channel 2 ",data)
				counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <- channel1:
				fmt.Println("Channel 1 ",data)
				counter++
			case data := <- channel2:
				fmt.Println("Channel 2 ",data)
				counter++
			default:
				fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}