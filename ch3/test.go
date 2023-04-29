package main

import "fmt"

func main() {
	var apples int32 = 123
	var oranges int64 = 23
	var a = int(apples) + int(oranges)
	fmt.Print(a)
}
