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
	fmt.Println("commencing to count down,press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:

		case <-abort:
			fmt.Println("launch aborted")
			return
		}
	}
	launch()
}
