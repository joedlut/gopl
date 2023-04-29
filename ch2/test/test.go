package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func incr(x *int) int {
	*x++
	return *x
}

func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}

func delta(old, new int) int {
	return new - old
}
func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(x)
	//false
	fmt.Println(f() == f())

	v := 1
	incr(&v)
	fmt.Println(v)
	fmt.Println(incr(&v))
	fmt.Println(v)

}
