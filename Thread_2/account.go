package main

//멀티 스레드 사용시 한 메모리를 공유해서 사용하기 때문에 실행 시간차에 따라서 원하는 결과값이 정상적으로 나오지 않을수가 있다.

//각 객체를 생성하는 구간 급격하게 메모리를 공유해서 빠르게 연산하는 구간에 mutex 라는 함수를 이용해서 락을 걸어주어 해당 문제을 일시적으로 해결 할수 있다.

//단 사용법이 너무 어려워 해당 함수를 사용할 때에는 사전 로직 설계가 반드시 필요하다.

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
}

func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()
	return total
}

func RandomTransfer() {
	var sender, balance int

	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance() 
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer(){
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex:&sync.Mutex{}})
	}
	globalLock = &sync.Mutex{}

	PrintTotalBalance()

	for i := 0; i < 10; i ++{
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}	


