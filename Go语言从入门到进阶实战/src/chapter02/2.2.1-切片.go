package main

import "fmt"

func main() {
	// 声明字符串切片
	// var strList [] string


	// // 声明整型切片
	// var numList [] int

	// var numEmpty  = [] int{}

	// // 输出3个切片
	// fmt.Println(strList, numList, numEmpty)

	// fmt.Println(len(strList), len(numList), len(numEmpty))

	// fmt.Println(strList == nil)

	// fmt.Println(numList == nil)

	// fmt.Println(numEmpty == nil)

	// var numbers [] int

	// for i:=0;i<10;i++{
	// 	numbers =append(numbers,i)

	// 	fmt.Printf("%d %d %p\n",len(numbers),cap(numbers), numbers)
	// }

	// 添加多个元素
	var car []string
	car = append(car, "LINK%CO", "Sniper", "Monk")

	// 添加切片
	team := []string{"Pig", "Chicken"}

	car = append(car, team...)

	fmt.Println(car)


}
