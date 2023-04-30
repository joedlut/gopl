package main

import "fmt"

func main() {
	ages := make(map[string]int)
	//ages["tom"] = 12
	//ages["joe"] = 13
	// 非法操作
	//tomValueAddr := &ages["joe"]
	//fmt.Println(tomValueAddr)
	fmt.Println(ages == nil)

	var ages1 map[string]int
	fmt.Println(ages1 == nil)

}
