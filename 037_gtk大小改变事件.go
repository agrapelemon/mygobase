package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/037.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//"configure_event"	窗口大小改变时触发
	win.Connect("configure_event", func() {
		var w, h int
		w, h = win.GetSize()
		fmt.Printf("w = %v, h = %v\n", w, h)
	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
