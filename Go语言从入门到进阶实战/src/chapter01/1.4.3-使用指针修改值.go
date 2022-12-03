package main

import (
	"fmt"
)

func swap(a, b *int) {
	// 获取a指针的值
	t := *a

	// 把指针b的值，赋值给a指向的变量
	*a = *b

	// 将a指针的值赋值给b指针指向的变量
	*b = t

	// 错误示范
	b, a = a, b // 结果 x，y的值不会变。
}
func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

}
