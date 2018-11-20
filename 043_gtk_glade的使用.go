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
	builder.AddFromFile("gladefiles/043.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))
	b1 := (*gtk.Button)(unsafe.Pointer(builder.GetObject("button1")))
	b2 := (*gtk.Button)(unsafe.Pointer(builder.GetObject("button2")))

	//信号处理
	b1.Clicked(func() {
		fmt.Println("按钮1111111111被按下")
	})

	b2.Clicked(func() {
		fmt.Println("按钮2222222222被按下")
	})

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//显示控件，如果是通过glade添加的控件，Show即可显示所有
	//         如果是通过代码布局添加的控件，需要ShowAll才能显示所有
	win.Show()

	//主事件循环
	gtk.Main()
}
