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
	builder.AddFromFile("gladefiles/035.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//添加鼠标按下事件
	//BUTTON_PRESS_MASK: 鼠标按下，触发信号"button-press-event"
	//BUTTON_MOTION_MASK: 鼠标移动，按下任何键移动都可以
	//BUTTON1_MOTION_MASK：鼠标移动，按住左键移动才触发
	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	//"button-press-event"	鼠标按下时触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		//1为左键，2为中间键，3为右键
		//fmt.Println(event.Button)
		flag := event.Button
		if flag == 1 {
			fmt.Println("1为左键")
		} else if flag == 2 {
			fmt.Println("2为中间键")
		} else if flag == 3 {
			fmt.Println("3为右键")
		}

		if event.Type == int(gdk.BUTTON_PRESS) {
			fmt.Println("单击")
		} else if event.Type == int(gdk.BUTTON2_PRESS) {
			fmt.Println("双击")
		}

		fmt.Printf("相对于窗口：%v, %v\n", event.X, event.Y)
		fmt.Printf("相对于屏幕：%v, %v\n", event.XRoot, event.YRoot)

	})

	//"motion-notify-event"	按住鼠标移动时触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		fmt.Printf("========相对于窗口：%v, %v\n", event.X, event.Y)
		fmt.Printf("========相对于屏幕：%v, %v\n", event.XRoot, event.YRoot)
	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
