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
	builder.AddFromFile("gladefiles/030.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取控件
	entry := (*gtk.Entry)(unsafe.Pointer(builder.GetObject("entry1")))

	//设置内容
	entry.SetText("123453253215")

	//获取内容
	fmt.Println(" entry.GetText() = ", entry.GetText())

	//是否只读
	//entry.SetEditable(false)

	//是否可见，密码模式
	//entry.SetVisibility(false)

	//"activate"	控件内部按回车键时触发
	entry.Connect("activate", func() {
		//获取内容
		fmt.Println(" entry.GetText() = ", entry.GetText())
	})

	//显示控件
	win.Show()

	//主事件循环
	gtk.Main()
}
