package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main() {
	srcImage := gocv.IMRead("/home/zdf/aaaaaa/gocvpicandvideo/1.jpg", 1)
	defer srcImage.Close()
	window := gocv.NewWindow("原图")
	//defer window.Close()
	window.IMShow(srcImage)

	dstImage := gocv.NewMat()
	defer dstImage.Close()
	var point image.Point
	point.X = 70
	point.Y = 70
	gocv.Blur(srcImage, &dstImage, point)
	win := gocv.NewWindow("腐蚀操作")
	//defer win.Close()
	win.IMShow(dstImage)
	gocv.WaitKey(0)
}
