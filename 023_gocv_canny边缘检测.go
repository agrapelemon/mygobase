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

	dstImage := gocv.NewMat()
	defer dstImage.Close()
	edge := gocv.NewMat()
	defer edge.Close()
	grayImage := srcImage.Clone()
	defer grayImage.Close()
	//dstImage.Reshape(srcImage.Cols(),srcImage.Rows())
	gocv.CvtColor(srcImage, &grayImage, gocv.ColorBGRToGray)
	var point image.Point
	point.X = 3
	point.Y = 3
	gocv.Blur(grayImage, &edge, point)
	gocv.Canny(edge, &edge, 3, 9)
	win := gocv.NewWindow("腐蚀操作")
	win.IMShow(edge)
	gocv.WaitKey(0)
}
