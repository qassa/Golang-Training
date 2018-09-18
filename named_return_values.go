package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum + 4*12
	y = sum / x
	return
}

func main() {
	x, y := split(2)
	fmt.Println(x, y)
	fmt.Println(split(4))
}
