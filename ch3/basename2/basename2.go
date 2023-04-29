package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}


func main() {
	fmt.Println(basename("/html/v2/3.html"))
	fmt.Println(basename("/html/v2/345"))
	fmt.Println(basename("/html/v2/."))

	teststr := "ddsfsdfsd"
	fmt.Println(len(teststr[:0]))

}
