package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
)

func GetURLs(name string) []string {
	defer file.Close()

	file, err := os.Open(name)
	if err != nil {
		panic("opening error")
	}
	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return []string{"", ""}
}

func WriteLines(urls []string) {
	for elem := range urls {

	}
}

func main() {
	var fileN string = os.Args[1]
	strings := GetURLs(fileN)
	fmt.Print(strings)
}
