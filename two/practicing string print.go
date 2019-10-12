package main

import "fmt"

var x = 42
var y = "Ross O'Brien"
var z = true

func main() {

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s := fmt.Sprintf("%v\t%v\t%v\t",x,y,z)

	fmt.Println(s)
}