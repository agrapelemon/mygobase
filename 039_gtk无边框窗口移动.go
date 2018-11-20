package main

import (
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/039.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))
	win.SetDecorated(false) //去边框

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//添加鼠标按下事件
	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	x, y := 0, 0

	//"button-press-event"	鼠标按下时触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		x, y = int(event.X), int(event.Y)

	})

	//"motion-notify-event"	按住鼠标移动时触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		win.Move(int(event.XRoot)-x, int(event.YRoot)-y)

	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
