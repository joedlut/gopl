package main

import (
	"fmt"
	"sync"
	"time"
)

type Rocket struct {
}

func (r *Rocket) Launch() { fmt.Println("rocket launch") }

func main() {
	var n sync.WaitGroup
	r := new(Rocket)
	n.Add(1)
	go func() {
		defer n.Done()
		time.AfterFunc(3*time.Second, r.Launch)
	}()
	n.Wait()
}
