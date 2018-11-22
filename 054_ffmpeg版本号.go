package main

/*
#cgo CFLAGS: -I/usr/local/ffmpeg/include
#cgo LDFLAGS: -L/usr/local/ffmpeg/lib -lavformat
#include "libavformat/avformat.h"
#include "libavcodec/avcodec.h"
#include "libavutil/avutil.h"
#include "libavutil/opt.h"
#include "libavdevice/avdevice.h"
*/
import "C"
import (
	"fmt"
)

func main() {

	fmt.Println(C.avformat_version())
	fmt.Printf("%T\n", C.avformat_version())

}
