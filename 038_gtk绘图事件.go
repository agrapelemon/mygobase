package main

import (
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/038.glade")

	//获取glade上的控件
	win := (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1")))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	var w, h int
	//"configure_event"	窗口大小改变时触发
	win.Connect("configure_event", func() {
		w, h = win.GetSize()
		//刷图，整个窗口区域刷图
		win.QueueDraw() //====> 触发"expose-event"
	})
	//允许窗口绘图
	win.SetAppPaintable(true)

	x := 0
	//绘图时所触发的信号：expose-event。
	win.Connect("expose-event", func() {
		//设置画家，指定绘图区域
		painter := win.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)

		//创建图片资源
		bg, _ := gdkpixbuf.NewPixbufFromFileAtScale("./Image/bk.jpg", w, h, false)
		face, _ := gdkpixbuf.NewPixbufFromFileAtScale("./Image/face.png", 80, 80, false)

		painter.DrawPixbuf(gc, bg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, face, 0, 0, x, 150, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

		//释放图片资源
		bg.Unref()
		face.Unref()
	})

	//获取按钮
	button := (*gtk.Button)(unsafe.Pointer(builder.GetObject("button1")))

	button.Clicked(func() {
		x += 50

		if x >= w {
			x = 0
		}

		//刷图，整个窗口区域刷图
		win.QueueDraw() //====> 触发"expose-event"
	})

	//显示控件
	win.ShowAll()

	//主事件循环
	gtk.Main()
}
