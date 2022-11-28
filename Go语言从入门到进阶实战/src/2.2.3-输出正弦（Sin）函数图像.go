package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func gen_Sin() {
	// 图片大小
	const size = 300

	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		// sin 值的范围 0 - 2Pi 之间
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}

	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal("Create file failed \n", err)
	}
	png.Encode(file, pic)
	file.Close()

}
func main() {
	gen_Sin()
}
