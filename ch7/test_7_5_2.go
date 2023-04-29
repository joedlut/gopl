package main

import (
	"bytes"
	"io"
)

// const debug = true
const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	// debug=false时候， out(var buf) 是包含空指针的非空接口 ， if条件为true
	if out != nil {
		//panic
		out.Write([]byte("hello"))
	}
}
