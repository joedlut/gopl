package main

import (
	"fmt"
	"sync"
)

func main() {
	var n sync.WaitGroup
	n.Add(1)
	go func() {
		defer n.Done()
		fmt.Print(0)
	}()
	fmt.Print(1)
	n.Wait()
}
