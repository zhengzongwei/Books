package main

import (
	"fmt"
)

func main() {
	var house = "Malibu Point 10086"
	ptr := &house

	// 指针类型
	fmt.Printf("ptr type %T\n", ptr)
	// 指针地址
	fmt.Printf("ptr %p\n", ptr)

	// 指针取值
	value := *ptr

	fmt.Printf("value type %T\n", value)

	fmt.Printf("value %s\n", value)

}

"""
ptr type *string
ptr 0xc000010250
value type string
value Malibu Point 10086
"""
