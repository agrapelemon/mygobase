package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func HandleSignal(ctx *glib.CallbackContext) {
	fmt.Println("===================")
	arg := ctx.Data()        //获取用户传递的参数，它是空接口类型
	data, ok := arg.(string) //类型断言
	if ok {
		fmt.Printf("按钮1被按下： %s\n", data)
	}
}

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

	//8、信号处理
	//按钮按下触发的信号："clicked"

	str := "are u ok?"
	//告诉系统，只要按下按钮，自动调用HandleSignal, str是给HandleSignal传递的参数
	//Connect()只会调用一次，告诉系统一个规则，只要告诉一次就够
	b1.Connect("clicked", HandleSignal, str)

	/*
		//处理函数可以是匿名函数，推荐写法
		b2.Connect("clicked", func() {
			fmt.Printf("按钮222222222被按下： %s\n", str)
		})
	*/

	//这种写法和上面等价
	b2.Clicked(func() {
		fmt.Printf("按钮222222222被按下： %s\n", str)
	})

	//按窗口的关闭按钮，触发"destroy"
	win.Connect("destroy", func() {
		fmt.Println("++++++++++++++++")

		gtk.MainQuit() //gtk程序关闭

	})

	//主事件循环（固定）
	gtk.Main()
	fmt.Println("窗口已关闭")

}
