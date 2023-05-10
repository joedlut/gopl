package bank3

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	//balance += amount
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	//defer mu.Unlock()
	return b
}

func WithDraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	//Deposit(-amount)
	deposit(amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}
