package main

import "fmt"

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0] //共用底层数组
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func noempty3(strings []string) []string {
	out := make([]string, 0)
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)

	data1 := []string{"1", "2", "", "5", ""}
	fmt.Printf("%q\n", noempty2(data1))
	fmt.Printf("%q\n", data1)

	data3 := []string{"1", "2", "", "6", ""}
	fmt.Printf("%q\n", noempty3(data3))
	fmt.Printf("%q\n", data3)

}
