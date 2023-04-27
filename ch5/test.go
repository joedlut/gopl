package main

import (
	"fmt"
	"strings"
)

func main() {
	newstr := strings.Map(func(r rune) rune {
		return r + 1
	}, "hello world")
	fmt.Println(newstr)
}
