package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture("/home/zdf/aaaaaa/gocvpicandvideo/1.mp4")
	window := gocv.NewWindow("显示视频")
	frame := gocv.NewMat()
	edges := gocv.NewMat()
	var point image.Point
	point.X = 7
	point.Y = 7
	for {
		if webcam.Read(&frame) {
			//window.IMShow(frame)
			gocv.CvtColor(frame, &edges, gocv.ColorBGRToGray)
			gocv.Blur(edges, &edges, point)
			gocv.Canny(edges, &edges, 0, 30)
			window.IMShow(edges)
			window.WaitKey(10)
		} else {
			window.WaitKey(0)
			break
		}
	}
}
