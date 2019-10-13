package main

import "fmt"

func main() {

	x := 42

	if x == 40 {
		fmt.Printf("The correct answer is: %d",x)
	} else if x == 12 {fmt.Printf("The correct answer is actually: %d",x)

	} else if x == 42{
		fmt.Printf("The real answer is %d",x)
	}

}
