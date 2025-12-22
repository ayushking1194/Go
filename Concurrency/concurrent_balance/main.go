package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu sync.Mutex
}

func (b *BankAccount) Deposit(amount int, wg *sync.WaitGroup){
	defer wg.Done()

	b.mu.Lock()
	defer b.mu.Unlock()

	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int, wg *sync.WaitGroup){
	defer wg.Done()

	b.mu.Lock()
	defer b.mu.Unlock()

	b.balance -= amount
}

func main() {
	account := BankAccount{balance: 100}

	var wg sync.WaitGroup
	wg.Add(2)

	go account.Deposit(50, &wg)
	go account.Withdraw(30, &wg)

	wg.Wait()

	fmt.Println("Final balance :-", account.balance)
}