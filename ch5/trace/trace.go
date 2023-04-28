package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	//延迟执行的是 trace返回的函数，语句后面的()不要忘了
	defer trace("big slow operation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s  %s s", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
