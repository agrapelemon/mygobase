package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	mat := gocv.NewMatWithSize(480, 640, gocv.MatTypeCV8UC4)
	//mat := gocv.IMRead("/home/zdf/aaaaaa/gocvpicandvideo/1.jpg", 1)
	defer mat.Close()
	createAlphaMat(&mat)
	fmt.Printf("%T\n", mat)
	gocv.IMWrite("010/1.png", mat)
	window := gocv.NewWindow("生成的PNG图")
	window.IMShow(mat)
	gocv.WaitKey(0)

}
func createAlphaMat(mat *gocv.Mat) {
	//m := 0xFF0000
	//fmt.Printf("m = %T\n", m)
	for i := 0; i < mat.Rows(); i++ {
		for j := 0; j < mat.Cols(); j++ {
			k := 0xFF000000
			k += i*256 + j
			fmt.Println(int32(k))
			mat.SetIntAt(i, j, int32(k))
		}
	}
}
