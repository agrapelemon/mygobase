package main

import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/029.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取按钮控件
	b1 := (*gtk.Button)(unsafe.Pointer(builder.GetObject("button1")))
	b2 := (*gtk.Button)(unsafe.Pointer(builder.GetObject("button2")))

	b1.SetLabel("^_^") //设置文本内容
	//b1.SetLabelFontSize(30)
	b1.ModifyFontEasy("DejaVu Serif 30")     //改变字体大小
	fmt.Println("b1 text = ", b1.GetLabel()) //获取内容
	b1.SetSensitive(false)                   //变灰，不让按

	//获取b2大小
	var w, h int
	w, h = b2.GetSizeRequest()

	//新建图片资源， 大小和b2差不多
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./Image/face.png", w-10, h-10, false)

	//新建image
	image := gtk.NewImageFromPixbuf(pixbuf)
	pixbuf.Unref() //释放资源
	//b2.SetImage(image) //按钮设置图标
	fmt.Println("按钮设置图片失败", image)

	b2.SetCanFocus(false) //取消焦距

	//显示控件
	win.Show()

	//主事件循环
	gtk.Main()
}
