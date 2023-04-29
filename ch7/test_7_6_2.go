package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 2, 1, 3, 4, 5}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)

	fmt.Println(sort.IntsAreSorted(values))

	// 先用sort.IntSlice 强转， values没有实现 Len, Less, Swap方法
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)

	fmt.Println(sort.IntsAreSorted(values))
}
