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
	r.Launch()
	n.Add(1)
	go func() {
		defer n.Done()
		fmt.Println(3)
		time.AfterFunc(3*time.Second, r.Launch)
	}()
	n.Add(1)
	go func() {
		defer n.Done()
		fmt.Println(5)
		time.AfterFunc(5*time.Second, r.Launch)
	}()
	//go time.AfterFunc(3*time.Second, r.Launch)
	n.Wait()
}
