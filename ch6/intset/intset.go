package main

import "fmt"

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func main() {
	i := IntSet{[]uint64{12, 34, 567}}
	fmt.Println(i.Has(23))
	fmt.Println(i.Has(34))
}
