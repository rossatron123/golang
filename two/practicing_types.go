package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n",x)
	y = int(x)
	fmt.Printf("%T\n",x)

}
