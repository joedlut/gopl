package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("launch success")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):

	case <-abort:
		fmt.Println("launch aborted")
		return
	}
	launch()
}
