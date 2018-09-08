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

	if texts, err := output(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		for _, text := range texts {
			fmt.Println(text)
		}
	}
}

func output(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	var texts []string
	for scan.Scan() {
		text := scan.Text()
		texts = append(texts, text)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return texts, nil
}
