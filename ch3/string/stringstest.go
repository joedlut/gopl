package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "hello,世界"
	charCount := utf8.RuneCountInString(str1)
	fmt.Println(charCount)

	for i := 0; i < len(str1); {
		r, size := utf8.DecodeRuneInString(str1[i:])
		fmt.Printf("%d \t %c\n", i, r)
		i += size
	}

	n := 0
	for range str1 {
		n++
	}
	fmt.Println(n)

	for i, r := range str1 {
		fmt.Printf("%d \t %c\n", i, r)
	}

	runeSlice := []rune(str1)
	fmt.Println(runeSlice)
}
