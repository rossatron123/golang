package main

import "fmt"

func main() {
	switch{
	case false:
		fmt.Println("This is false")
	case true:
		fmt.Println("This is true")
		fallthrough
	case false:
		fmt.Println("This prints because it's a fallthrough")

	}
}
