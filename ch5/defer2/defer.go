package main

import (
	"os"
	"runtime"
)

func f(x int) {
	defer printStack()
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	f(10)

	recover()
}
