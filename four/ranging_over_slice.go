package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[2])
	for i, v := range x{
		fmt.Println(i,v)
	}

}
