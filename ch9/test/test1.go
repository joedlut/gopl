package main

import "fmt"

func main() {
	for {
		fmt.Print(1)
		go fmt.Print(0)
	}
}
