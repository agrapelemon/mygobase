package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	if true {
		fmt.Println("--1.zlib普通压缩---------------")
		var in bytes.Buffer
		b := []byte(`12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890`)
		fmt.Println("原文：", string(b))
		w := zlib.NewWriter(&in)
		w.Write(b)
		w.Close()
		fmt.Println(in.Bytes())
		fmt.Println("压缩后长度：", len(in.Bytes()))

		var out bytes.Buffer
		r, _ := zlib.NewReader(&in)
		io.Copy(&out, r)

		fmt.Println("原文：", out.String())
		fmt.Println("解压后长度：", len(out.Bytes()))
	}
}
