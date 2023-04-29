package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("%T\n", f)
	//panic
	//c := w.(*bytes.Buffer)
	//fmt.Printf("%T", c)

	if f, ok := w.(*os.File); ok {
		fmt.Printf("%T\n", f)
		fmt.Println(f)
	}
}
