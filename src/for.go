package main

import "fmt"

func main() {
	var abc []string
	abc = []string{"a", "b", "c"}

	for i, v := range abc {
		fmt.Println(i, v)
	}
	fmt.Println()

	for _, v := range abc {
		fmt.Println(v)
	}
	fmt.Println()

	for i := range abc {
		fmt.Println(i)
	}
}
