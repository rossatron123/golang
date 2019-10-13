package main

import "fmt"

func main() {

	m := map[string][]string{

		"Ross": []string{"One","Two","Three"},
		"Rachel": []string{"Ask","The","Question"},
	}

	fmt.Println(m)

	for k, v := range m{
		fmt.Println("The person named:",k)
		for i, v2 := range v{
			fmt.Println("\t\t",i,v2)
		}
	}
}
