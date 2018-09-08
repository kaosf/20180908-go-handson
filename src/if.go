package main

import "fmt"

func main() {
	v := "GoLang"
	if v == "golang" {
		fmt.Println("○")
	} else if v == "GOLANG" {
		fmt.Println("△")
	} else {
		fmt.Println("☓")
	}

	fmt.Println(v)
}
