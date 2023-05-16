package main

import (
	"fmt"
	"time"
)

type Rocket1 struct {
}

func (r *Rocket1) Launch() {
	fmt.Println("rocket launch")
	ch <- 30
}

var ch = make(chan int)

func main() {

	r := new(Rocket1)
	time.AfterFunc(3*time.Second, r.Launch)
	for {
		select {
		case n := <-ch:
			fmt.Printf("%d arrived\n", n)
			fmt.Println("done")
			return

		default:
			fmt.Println("time to wait")
			time.Sleep(3 * time.Second)
		}
	}

}
