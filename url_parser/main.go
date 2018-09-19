package main

import (
	"fmt"
	"os"
	"parser"
)

func main() {
	fname := os.Args[1]
	strings, err := parser.ParseFile(fname)
	if err != nil {
		panic(err)
	}
	var lens []int
	for i, uri := range strings {
		lens = append(lens, parser.GetAndWrite(uri, i+1))
	}
	for i, uri := range strings {
		fmt.Printf("%d. %s\n", i+1, uri)
		fmt.Printf("the length of page body is %d symbols\n", lens[i])
	}

}
