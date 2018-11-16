package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	if true {
		fmt.Println("--1.显示图片---------------")
		img := gocv.IMRead("/home/zdf/aaaaaa/gocvpicandvideo/1.jpg", 199)
		window := gocv.NewWindow("显示图片")
		window.IMShow(img)
		window.WaitKey(2000)
	}
	if true {
		fmt.Println("--2.显示视频---------------")
		webcam, _ := gocv.OpenVideoCapture("/home/zdf/aaaaaa/gocvpicandvideo/1.mp4")
		window := gocv.NewWindow("显示视频")
		img := gocv.NewMat()
		for {
			if webcam.Read(&img) {
				window.IMShow(img)
				window.WaitKey(1)
			} else {
				window.WaitKey(0)
				break
			}
		}
	}

}
