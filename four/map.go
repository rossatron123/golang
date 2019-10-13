package main

import "fmt"

func main() {
	m := map[string]int{

		"Ross": 26,
		"O'Brien": 6,
	}
	fmt.Println(m)

	if v, ok := m["Ross"]; ok{
		fmt.Println("The age of Ross is",v)
	}
}
