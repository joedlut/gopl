package bank4

import "sync"

var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
