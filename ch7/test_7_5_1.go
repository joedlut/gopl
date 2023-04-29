package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	//var x interface{} = time.Now()
	//fmt.Println(x)

	//panic
	//var x interface{} = []int{1, 2, 3}
	//fmt.Println(x == x)

	var x io.Writer
	fmt.Printf("%T\n", x)

	x = os.Stdout
	fmt.Printf("%T\n", x)

	x = new(bytes.Buffer)
	fmt.Printf("%T\n", x)

}
