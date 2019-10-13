package main

import "fmt"

func main() {
	m := map[string]int{

		"Ross":26,
		"Jason":34,
		"Rachel":30,
	}

	fmt.Println(m)

	m["Mum"] = 60
	fmt.Println(m)

	delete(m,"Jason")
	fmt.Println(m)

	if v, ok := m["Mum"]; ok{
		fmt.Println("Mum is here!",v)
	}

	delete(m,"Rachel")
	fmt.Println(m)
}
