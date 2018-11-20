package main

import (
	"fmt"
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
	builder.AddFromFile("gladefiles/036.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//"key-press-event"	键盘按下时触发
	win.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		//获取键盘按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventKey)(unsafe.Pointer(&arg))

		//获取到实际上是字母的ascii
		key := event.Keyval
		//fmt.Printf("key = %v\n", key)
		//if key == 97 {
		if key == gdk.KEY_a {
			fmt.Println("aaaaaaaaaaaaa")
		}
	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
