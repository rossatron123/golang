package main

import "fmt"

func main() {
	for i := 0; i < 10; i++{
		for j := 0; j <3; j++{
			fmt.Printf("The outer number is %d, and the inner number is %d\n",i,j)
		}
	}
}
