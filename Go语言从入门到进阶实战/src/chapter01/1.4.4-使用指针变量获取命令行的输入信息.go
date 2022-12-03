package main

import (
	"flag"
	"fmt"
)



var mode = flag.String("mode","","process mode")
func main()  {
	
	flag.Parse()

	fmt.Println(*mode)
}

// go run 1.4.4-使用指针变量获取命令行的输入信息.go --mode=kkk