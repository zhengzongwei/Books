package main

import (
	"fmt"
	"unicode/utf8"
)


func main(){

	tip1 := "this is a ninja"
	fmt.Println(len(tip1)) // 15

	tip2 := "你好，世界"
	fmt.Println(len(tip2))  // 15
	fmt.Println(utf8.RuneCountInString(tip1)) // 15
	fmt.Println(utf8.RuneCountInString(tip2)) // 5
}

