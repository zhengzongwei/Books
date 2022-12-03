package main

import (
	"fmt"
)

func main() {
	theme := "你好 World"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascill %c %d\n", theme[i], theme[i])
	}
	/*
		ascill ä 228
		ascill ½ 189
		ascill   160
		ascill å 229
		ascill ¥ 165
		ascill ½ 189
		ascill   32
		ascill W 87
		ascill o 111
		ascill r 114
		ascill l 108
		ascill d 100
	*/

	for _, s := range theme {
		fmt.Printf("Unicode %c %d\n", s, s)
	}
	/*
		Unicode 你 20320
		Unicode 好 22909
		Unicode   32
		Unicode W 87
		Unicode o 111
		Unicode r 114
		Unicode l 108
		Unicode d 100
	*/
}
