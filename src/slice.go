package main

import "fmt"

func main() {
	var example []string
	fmt.Println("初期値は", example)

	example = []string{
		"Golang",
		"hands-on",
	}
	fmt.Println("2番目の値は", example[1])

	example = append(example, "in", "kagawa")
	fmt.Println("追加結果は", example)

	fmt.Println(example[0:2])
	fmt.Println(example[2:])

	example[3] = "okayama"
	fmt.Println(example)
}
