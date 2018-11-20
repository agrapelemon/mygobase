package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//1、初始化（固定）
	gtk.Init(&os.Args)

	//用户写的代码
	//a)创建窗口
	//b)设置属性(标题，大小)
	//c) 显示窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("go gtk")                    //设置标题
	win.SetSizeRequest(480, 320)              //设置大小
	win.Show()                                //显示

	fmt.Println("before")
	//3、主事件循环（固定）
	//a) 让程序不结束 b) 等待用户操作（移动窗口、点击鼠标）
	gtk.Main()
	fmt.Println("over")

}
