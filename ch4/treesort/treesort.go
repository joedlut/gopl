package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 并不是平衡二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	var t *tree
	v := []int{12, 45, 67, 8, 89, 23, 1, -90, 23, -1}
	var result []int
	for _, n := range v {
		t = add(t, n)
	}
	result = appendValues(result, t)
	fmt.Println(result)
	//test
}
