package main

import "os"

func main() {
	gtk.Init(&os.Args) //环境初始化

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	window.SetTitle("GTK Go!")                   //设置标题
	window.SetSizeRequest(300, 200)              //设置窗口的宽度和高度

	window.Show() //显示窗口

	gtk.Main() //主事件循环，等待用户操作

}
