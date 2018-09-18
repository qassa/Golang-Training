package main

import "fmt"

var c, python, java = true, false, "no!"

var i, j = 1, 2
var a, b int = 1, 2

//var i int, j int = 1, 2 - не работает, т.к. var определяет список однотипных переменных
//var i int, j int - не работает по той же причине

func main() {
	var d int
	fmt.Println(d, c, python, java)
	fmt.Println(i, j, a, b)
}
