package main

import "fmt"

func main() {
	fmt.Print("入力は？>")
	var ans string

	fmt.Scanln(&ans)

	fmt.Printf("入力は %s です\n", ans)
}
