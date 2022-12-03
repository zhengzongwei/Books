package main

import "fmt"

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}
func main() {
	fmt.Println(calc(3, 4))
}
