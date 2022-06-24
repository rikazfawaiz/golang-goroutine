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

// Test Deadlock, deadlock keadaan dimana dua atau lebih gourutine nge-lock secara bersamaan saling tunggu
type UserBalance struct{
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("Lock User1",user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2",user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T){
	user1 := UserBalance{
		Name : "Rikaz",
		Balance : 1000000,
	}

	user2 := UserBalance {
		Name : "Fawaiz",
		Balance : 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User",user1.Name,", Balance",user1.Balance)
	fmt.Println("User",user2.Name,", Balance",user2.Balance)

	
}