package main

import (
	"fmt"
)

func main() {
	x := 42
	y := 100
	fmt.Println(x)
	fmt.Println(y)

	x = 6
	fmt.Println(x)

	y = y + x
	fmt.Println(y)
}