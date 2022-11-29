package main

import "fmt"


func main()  {
	var a byte = 'a'
	fmt.Printf("%d %T\n", a, a)

	var b rune = '你'
	fmt.Printf("%d %T\n", b, b)
	
}

// 结果
// 97 uint8
// 20320 int32