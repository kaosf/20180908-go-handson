package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid args")
		os.Exit(1)
	}
	filename := os.Args[1]
	if _, err := os.Stat(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		fmt.Println(text)
	}
	if err := scan.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
