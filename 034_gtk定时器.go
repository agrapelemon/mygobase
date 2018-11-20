package main

import (
	"os"
	"strconv"
	"unsafe"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/034.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取按钮
	buttonStart := (*gtk.Button)(unsafe.Pointer(builder.GetObject("buttonStart")))
	buttonStop := (*gtk.Button)(unsafe.Pointer(builder.GetObject("buttonStop")))

	//获取标签
	label := (*gtk.Label)(unsafe.Pointer(builder.GetObject("label1")))
	label.SetText("0")                      //设置内容
	label.ModifyFontEasy("DejaVu Serif 30") //改变字体大小

	buttonStop.SetSensitive(false) //按钮变灰

	num := 0
	//id := 0

	btemp := false
	glib.TimeoutAdd(500, func() bool {
		if btemp {

			num++

			//strconv.Itoa(num)整型转字符串
			label.SetText(strconv.Itoa(num))

			return true //注意，要返回ture
		} else {
			return true
		}
	})
	//启动定时器
	buttonStart.Clicked(func() {
		btemp = true
		buttonStart.SetSensitive(false)
		buttonStop.SetSensitive(true)
	})

	//停止定时器
	buttonStop.Clicked(func() {
		btemp = false
		buttonStart.SetSensitive(true)
		buttonStop.SetSensitive(false)

	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
