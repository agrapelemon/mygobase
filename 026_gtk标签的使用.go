package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gtk"

	"github.com/mattn/go-gtk/gdk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/026.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序d
	})

	//获取标签控件
	labelOne := (*gtk.Label)(unsafe.Pointer(builder.GetObject("label1")))
	labelTwo := (*gtk.Label)(unsafe.Pointer(builder.GetObject("label2")))

	labelOne.SetText("你大爷")                                    //设置内容
	labelOne.ModifyFontEasy("DejaVu Serif 30")                 //设置字体大小
	labelOne.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white")) //设置字体颜色

	labelTwo.SetText("are u ok?")
	//获取内容
	str := labelTwo.GetText()
	fmt.Println("str = ", str)

	//显示控件，如果是通过glade添加的控件，Show即可显示所有
	//         如果是通过代码布局添加的控件，需要ShowAll才能显示所有
	win.Show()

	//主事件循环
	gtk.Main()
}

//package main
//
//import (
//	"os"
//
//	"fmt"
//	"github.com/mattn/go-gtk/gdk"
//	"github.com/mattn/go-gtk/gtk"
//	"unsafe"
//)
//
//func main() {
//	gtk.Init(&os.Args) //环境初始化
//
//	//window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
//	//window.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
//	//window.SetTitle("GTK Go!")                   //设置标题
//	//window.SetSizeRequest(300, 200)              //设置窗口的宽度和高度
//	//window.Show() //显示窗口
//	//fmt.Printf("%T\n", window)
//
//	builder := gtk.NewBuilder()
//	builder.AddFromFile("gladefiles/026.glade")
//	win := gtk.WidgetFromObject(builder.GetObject("window1"))
//
//	labelOne := (*gtk.Label)(unsafe.Pointer(builder.GetObject("label1")))
//	labelTwo := (*gtk.Label)(unsafe.Pointer(builder.GetObject("label2")))
//
//	//设置字体颜色
//	labelOne.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("green"))
//
//	//设置内容
//	labelOne.SetText("你大爷")
//
//	//设置字体大小
//	labelOne.ModifyFontEasy("DejaVu Serif 30")
//
//	//设置控件大小
//	labelOne.SetSizeRequest(200, 100)
//
//	//不起作用
//	labelOne.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("red"))
//
//	labelTwo.SetText("Are u ok ?")
//	//获取内容
//	str := labelTwo.GetText()
//	fmt.Println("str = ", str)
//	win.Show()
//	gtk.Main() //主事件循环，等待用户操作
//}
