package main

import "fmt"

func double(x int) (result int) {
	//末尾的()说明传递的参数是空
	//defer 语句在return执行后执行, 所以result 是8
	defer func() { fmt.Printf("double(%d)=%d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	_ = double(4)
	fmt.Println(triple(4))
}
