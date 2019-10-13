package main

import "fmt"

func main() {

	for i := 33; i <= 90; i++{
		fmt.Printf("The number is %v, and the hex is %#U \n",i,i)
	}

}
