package main

import (
	"fmt"
)

func main() {
	s := "Hello world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {

		fmt.Printf("%#U\n", s[i])
	}
}