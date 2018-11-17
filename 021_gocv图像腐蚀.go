package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main() {
	srcImage := gocv.IMRead("/home/zdf/aaaaaa/gocvpicandvideo/1.jpg", 1)
	defer srcImage.Close()
	window := gocv.NewWindow("原图")
	window.IMShow(srcImage)
	var point image.Point
	point.X = 15
	point.Y = 15
	element := gocv.GetStructuringElement(gocv.MorphRect, point)

	defer element.Close()
	dstImage := gocv.NewMat()
	defer dstImage.Close()
	gocv.Erode(srcImage, &dstImage, element)
	win := gocv.NewWindow("腐蚀操作")
	win.IMShow(dstImage)
	gocv.WaitKey(0)
}
