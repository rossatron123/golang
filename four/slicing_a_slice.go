package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}

	fmt.Println(x)
	fmt.Println(x[3])
	fmt.Println(x[2:4]) // includes what's at position 3 and up to but not including what is at position 5
}
