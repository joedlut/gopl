package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, i := range vals {
		total += i
	}
	return total
}

func sum1(vals ...int) {}
func sum2([]int)       {}
func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(23, 4))

	fmt.Printf("%T\n", sum1)
	fmt.Printf("%T\n", sum2)
}
