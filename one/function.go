package main

import "fmt"

func main() {
	fmt.Println("This includes a function")

	for i := 0; i < 100; i++ {
		if i%2 == 0{
			fmt.Println(i)
		}

	}
	foo()
}

func foo(){
	fmt.Println("This is the function foo")
}

