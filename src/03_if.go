package main

import "fmt"

var (
	example = []string{"golang", "hands-on", "in", "kagawa"}
)

func main() {
	var datas []string
	datas = example
	for _, v := range datas {
		fmt.Println(v)
		if v == "in" {
			fmt.Println("○")
		} else {
			fmt.Println("☓")
		}
	}
}
