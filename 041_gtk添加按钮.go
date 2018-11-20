package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//1、创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

	//2、设置窗口属性(大小、标题)
	win.SetTitle("呵呵呵")
	win.SetSizeRequest(480, 320)

	//3、创建容器控件（固定布局，任意布局）
	layout := gtk.NewFixed()

	//4、布局添加到窗口上
	win.Add(layout)

	//5、创建按钮
	b1 := gtk.NewButtonWithLabel("^_^")
	b2 := gtk.NewButtonWithLabel("@_@")
	b2.SetSizeRequest(100, 100) //设置按钮2大小

	//6、按钮添加到布局中
	layout.Put(b1, 0, 0)
	layout.Put(b2, 100, 100)

	//7、显示控件
	//如果有多个控件，如果使用Show, 需要每一个控件都要Show
	//	win.Show()
	//	layout.Show()
	//	b1.Show()
	win.ShowAll() //所有控件都显示

	//主事件循环（固定）
	gtk.Main()

}
