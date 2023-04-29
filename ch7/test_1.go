package main

import (
	"fmt"
	"io"
)

type Writer interface {
	Write(p []byte) (int, error)
}

type Stringer interface {
	String() string
}

type ReadeWriter interface {
	//io.Reader
	Read(p []byte) (n int, err error)
	io.Writer
	io.Closer
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func main() {
	var c ByteCounter
	c.Write([]byte("hello,world"))
	fmt.Println(c)

	c = 0
	var name = "dolly1"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)
}
