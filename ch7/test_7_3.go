package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	//w = time.Second
	w = new(bytes.Buffer)
	fmt.Println(w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer)
	fmt.Println(rwc)

	var any interface{}
	any = true
	any = 1.23
	any = "joedlut"
}
