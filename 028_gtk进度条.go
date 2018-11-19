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
	builder.AddFromFile("gladefiles/028.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取进度条控件
	//pg := gtk.ProgressBarFromObject(builder.GetObject("progressbar1"))
	pg := (*gtk.ProgressBar)(unsafe.Pointer(builder.GetObject("progressbar1")))

	//设置进度 0.0 ~ 1.0
	pg.SetFraction(0.33)

	//获取进度，并打印
	fmt.Println("value = ", pg.GetFraction())

	//设置文本内容
	pg.SetText("33%")

	//显示控件
	win.Show()

	//主事件循环
	gtk.Main()
}
