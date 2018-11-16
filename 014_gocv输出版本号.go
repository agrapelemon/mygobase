package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	fmt.Println(gocv.Version())
	fmt.Println(gocv.OpenCVVersion())
}
