package main

import "fmt"

func main() {
	// 声明字符串切片
	var strList [] string


	// 声明整型切片
	var numList [] int

	var numEmpty  = [] int{}

	// 输出3个切片
	fmt.Println(strList, numList, numEmpty)

	fmt.Println(len(strList), len(numList), len(numEmpty))

	fmt.Println(strList == nil)

	fmt.Println(numList == nil)

	fmt.Println(numEmpty == nil)



	


}
