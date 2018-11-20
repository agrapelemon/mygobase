package main

import (
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/032.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取控件
	//layout := gtk.TableFromObject(builder.GetObject("table1"))
	layout := (*gtk.Table)(unsafe.Pointer(builder.GetObject("table1")))

	//新建一个按钮
	button := gtk.NewButtonWithLabel("新按钮")
	layout.AttachDefaults(button, 0, 1, 1, 3)

	button2 := gtk.NewButtonWithLabel("新按钮2")
	layout.AttachDefaults(button2, 2, 3, 0, 1)

	//通过代码添加的空间，需要显示所有控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
