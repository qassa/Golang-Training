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
	c := make(chan func() (int, int))
	for i, uri := range strings {
		go parser.GetAndWrite(uri, i+1, c)
	}
	for j, _ := range strings {
		_ = j
		i, len := (<-c)()
		fmt.Printf("%d. %d symbols is the length of %s\n", i, len, strings[i-1])
	}

}
