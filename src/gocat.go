package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"
	f, _ := os.Open(filename)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		fmt.Println(text)
	}
}
