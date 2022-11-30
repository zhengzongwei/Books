package main

import (
	"fmt"
)


func main()  {
	var cat int = 1
	var str string = "banana"

	fmt.Printf("%p %p", &cat, &str)
}

// 0xc0000ba000 0xc00009e210