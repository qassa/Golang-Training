package main

import(
  "fmt"
)

func add(x int, y int) int {
  return x + y
}

func swap(x, y string) (string, string){
return y, x
}

func main(){
  fmt.Println(46+20)
  a,b := swap("hello", "world")
  fmt. Println(a, b)
}