package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

//Handle Race Condition
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i:=1; i<=1000; i++ {
		go func(){
			for j:=1; i<=100; j++ {
				mutex.Lock() // agar aman ketika share variabel ke goroutine lain dan gak kena race condition
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ",x)
}

type BankAccount struct {
	RWMutex sync.RWMutex //Read Write Mutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // Write
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // Read
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i:=0; i<100; i++ {
		go func(){
			for j:=0; j<100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())

			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance",account.GetBalance())
}