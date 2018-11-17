package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	srcImage := gocv.IMRead("/home/zdf/aaaaaa/gocvpicandvideo/1.jpg", 1)
	defer srcImage.Close()
	dstImage := gocv.NewMat()
	defer dstImage.Close()
	fmt.Printf("%T\n", srcImage)
	fmt.Printf("%T\n", dstImage)
	gocv.CvtColor(srcImage, &dstImage, gocv.ColorBGRToLab)
	window := gocv.NewWindow("原图")
	window.IMShow(srcImage)
	window = gocv.NewWindow("颜色转换")
	window.IMShow(dstImage)
	gocv.WaitKey(0)

}
