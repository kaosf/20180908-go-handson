package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("golang ")
	fmt.Print("2018\n")

	fmt.Printf("%s %d\n", "golang", 2018)

	fmt.Println("golang", "2018")

	fmt.Fprintln(os.Stderr, "golang", "2018")

	s := fmt.Sprintf("%s %d\n", "golang", 2018)
	fmt.Println(s)
}
