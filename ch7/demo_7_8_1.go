package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))

	//类型转换
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
