package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {

	if true {
		fmt.Println("是否是绝对路径：", path.IsAbs("/aa"))
	}
	if true {
		dir, file := path.Split("aa/bb/cc/dd")
		fmt.Println("目录：", dir, "，文件：", file)
	}
	if true {
		file := path.Join("/aa", "bb")
		fmt.Println("文件：", file)
	}
	if true {
		file := path.Dir("/aa")
		fmt.Println("目录：", file)
	}
	if true {
		file, _ := filepath.Rel("aa", "aa/bb")
		fmt.Println("绝对路径转化成相对路径：", file)
	}
	if true {
		file := path.Clean("/tmp/a")
		fmt.Println("clean目录：", file)
	}
	if true {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Println(dir)

	}
	if true {
		dir, _ := filepath.Abs(".")
		fmt.Println("相对路径变成绝对路径：", dir)
	}
	if true {
		file := path.Ext("/tmp/a.1.avi")
		fmt.Println("扩展名：", file)
	}
}
