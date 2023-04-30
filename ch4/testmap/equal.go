package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		//0
		//fmt.Println(y[k])
		//错误的写法  直接 y[k] != xv  0==0
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 40}))
	fmt.Println(equal(map[string]int{"B": 40}, map[string]int{"B": 40}))
}
