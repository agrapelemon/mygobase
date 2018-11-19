package main

import (
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
	builder.AddFromFile("gladefiles/027.glade")

	//获取glade上的控件
	win := gtk.WidgetFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	//获取图片控件
	image := (*gtk.Image)(unsafe.Pointer(builder.GetObject("image1")))

	//获取控件大小
	var w, h int
	w, h = image.GetSizeRequest()

	//设置一张图片资源， pixbuf，控件大小和图片大小一样
	//false： 不保存图片原来的尺寸
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./Image/Luffy.png", w-10, h-10, false)

	image.SetFromPixbuf(pixbuf) //给image设置图片

	//图片资源使用完毕，需要释放空间（注意）
	pixbuf.Unref()

	//显示控件，如果是通过glade添加的控件，Show即可显示所有
	//         如果是通过代码布局添加的控件，需要ShowAll才能显示所有
	win.Show()

	//主事件循环
	gtk.Main()
}
