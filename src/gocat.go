package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		fmt.Println(text)
	}
	if err := scan.Err(); err != nil {
		return
	}
}
