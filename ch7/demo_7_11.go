package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Printf("%#v\n", err)

	fmt.Println(os.IsNotExist(err))

}
