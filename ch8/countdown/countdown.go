package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("launch success")
}

func main() {
	fmt.Println("commencing to count down")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
