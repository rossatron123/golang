package main

import "fmt"

func main() {
	x := make([]int,10,100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[2] = 42
	fmt.Println(x)
}
