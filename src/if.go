package main

import "fmt"

func main() {
	if v := "GoLang"; v == "golang" {
		fmt.Println("○")
	} else if v == "GOLANG" {
		fmt.Println("△")
	} else {
		fmt.Println("☓")
	}

	// fmt.Println(v)
}
