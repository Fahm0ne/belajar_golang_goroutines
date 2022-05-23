package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	 x := 0
	var mutex sync.Mutex 

	for i:=1; i <= 1000; i++ {
		go func () {
	for j:=1; j <= 100; j++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Perulangen ke = " , x)

}

type BankAccount struct {
Rwm sync.RWMutex
Balance int
}

func (b *BankAccount) AddBalance (amount int) {
	b.Rwm.Lock()
	b.Balance = b.Balance + amount
	b.Rwm.Unlock()

}

func (b *BankAccount) GetBalance () int {
	b.Rwm.RLock()
	balance := b.Balance
	b.Rwm.RUnlock()
	return balance
}

func TestRwmMutex(t *testing.T) {

	account :=	BankAccount{}

	for i:=0; i<100; i++ {
		go func ()  {
			for j:=0; j<100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
		fmt.Println("Total balance : " , account.GetBalance())	

}

type UserBalance struct {
	sync.Mutex
	user string
	balance int
}

func (u *UserBalance) ForLock () {
	u.Mutex.Lock()
}

// func (u *UserBalance) ForUnlock () {
// 	u.Mutex.Unlock()
// }

func (u *UserBalance)  Change (amount int) {
	u.balance = u.balance + amount
}

func Transfer (user1 *UserBalance , user2 *UserBalance , amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.user )
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2" , user2.user)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {

user1 := UserBalance {
	user : "Fahmi",
	balance: 10000,
}

user2 := UserBalance {
	user : "Firdaus",
	balance: 10000,
}

go Transfer(&user1 , &user2 , 1000)
go Transfer(&user2 , &user1 , 2000)


time.Sleep(3 * time.Second)

fmt.Println("User " , user1.user , "Balance ", user1.balance)
fmt.Println("User " , user2.user , "Balance ", user2.balance)


}

