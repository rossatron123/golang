package main

import "fmt"

var a int
type hotdog int
var b hotdog

func main() {
	a = 42
	b = 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)

}
