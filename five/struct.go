package main

import "fmt"

type person struct{

	first string
	last string
	age int
}

type secretagent struct{
	person
	ltk bool
}

func main() {

	p1 := person{

		first:"Ross",
		last: "O'Brien",
		age:26,
	}

	sa1 := secretagent{
		person:person{
			first: "Tom",
			last: "Smith",
			age: 12,
		},

		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(sa1)

}
