package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window //窗口
	buttonMin   *gtk.Button //最小化按钮
	buttonClose *gtk.Button //关闭按钮
	labelBlack  *gtk.Label  //记录黑棋个数
	labelWhite  *gtk.Label  //记录白棋个数
	labelTime   *gtk.Label  //记录倒计时
	imageBlack  *gtk.Image  //提示该黑子落子
	imageWhite  *gtk.Image  //提示该白子落子
}

//控件属性结构体
type ChessInfo struct {
	w, h           int //窗口的宽度和高度
	x, y           int //鼠标点击，相当于窗口的坐标
	startX, startY int //棋盘起点坐标
	gridW, gridH   int //棋盘一个格子的宽度和高度
}

//枚举，标志黑子白子状态
const (
	Empty = iota //当前棋盘格子没有子
	Black        //当前棋盘格子为黑子
	White        //当前棋盘格子为白子
)

//黑白棋结构体
type Chessboard struct {
	ChessWidget //匿名字段
	ChessInfo

	currentRole int //该谁落子
	tipTimerId  int //实现提示闪烁效果的定时器id
	leftTimerId int //倒计时定时器id
	timeNum     int //记录剩余时间

	chess [8][8]int //二维数组，标志棋盘状态
}

//函数：给按钮设置图标
func ButtonSetImageFromFile(button *gtk.Button, filename string) {
	//获取按钮的大小
	w, h := 0, 0
	w, h = button.GetSizeRequest()

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//创建image
	image := gtk.NewImageFromPixbuf(pixbuf)

	//释放pixbuf
	pixbuf.Unref()

	//给按钮设置图片
	button.SetImage(image)

	//去掉按钮的焦距
	button.SetCanFocus(false)
}

//给image设置图片
func ImageSetPicFromFile(image *gtk.Image, filename string) {
	//获取image的大小
	w, h := 0, 0
	w, h = image.GetSizeRequest()

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//给image设置图片
	image.SetFromPixbuf(pixbuf)

	//释放pixbuf
	pixbuf.Unref()
}

//方法： 创建控件，设置控件属性
func (obj *Chessboard) CreateWindow() {
	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("gladefiles/050.glade")

	//窗口相关
	obj.window = (*gtk.Window)(unsafe.Pointer(builder.GetObject("window1"))) //获取控件
	obj.window.SetAppPaintable(true)                                         //允许绘图
	obj.window.SetPosition(gtk.WIN_POS_CENTER)                               //居中显示
	obj.w, obj.h = 800, 480                                                  //窗口的宽度和高度
	obj.window.SetSizeRequest(800, 480)                                      //设置窗口的宽高
	obj.window.SetDecorated(false)                                           //去边框

	//设置事件，让窗口可以捕获鼠标点击和移动
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	//按钮相关
	//获取按钮控件
	obj.buttonMin = (*gtk.Button)(unsafe.Pointer(builder.GetObject("buttonMin")))
	obj.buttonClose = (*gtk.Button)(unsafe.Pointer(builder.GetObject("buttonClose")))

	//给按钮设置图片
	ButtonSetImageFromFile(obj.buttonMin, "./image/min.png")
	ButtonSetImageFromFile(obj.buttonClose, "./image/close.png")

	//标签相关
	obj.labelBlack = (*gtk.Label)(unsafe.Pointer(builder.GetObject("labelBlack")))
	obj.labelWhite = (*gtk.Label)(unsafe.Pointer(builder.GetObject("labelWhite")))
	obj.labelTime = (*gtk.Label)(unsafe.Pointer(builder.GetObject("labelTime")))

	//设置字体大小
	obj.labelBlack.ModifyFontEasy("DejaVu Serif 50")
	obj.labelWhite.ModifyFontEasy("DejaVu Serif 50")
	obj.labelTime.ModifyFontEasy("DejaVu Serif 30")
	//obj.labelBlack.ModifyFontSize(50)
	//obj.labelWhite.ModifyFontSize(50)
	//obj.labelTime.ModifyFontSize(30)

	//设置内容
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	obj.labelTime.SetText("20")

	//改变字体颜色
	obj.labelBlack.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelWhite.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelTime.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))

	//image相关
	obj.imageBlack = (*gtk.Image)(unsafe.Pointer(builder.GetObject("imageBlack")))
	obj.imageWhite = (*gtk.Image)(unsafe.Pointer(builder.GetObject("imageWhite")))

	//设置图片
	ImageSetPicFromFile(obj.imageBlack, "./image/black.png")
	ImageSetPicFromFile(obj.imageWhite, "./image/white.png")

	//棋盘相关
	obj.startX, obj.startY = 200, 60
	obj.gridW, obj.gridH = 50, 40
}

//方法：改变落子角色
func (obj *Chessboard) ChangeRole() {
	//重新设置时间
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))

	//先隐藏提示
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()

	//原来是黑子下，就到白子下
	if obj.currentRole == Black {
		obj.currentRole = White
	} else {
		obj.currentRole = Black
	}
}

//鼠标点击事件函数
func MousePressEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MousePressEvent Chessboard err")
		return
	}

	//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	//保存点击的x, y坐标
	obj.x, obj.y = int(event.X), int(event.Y)
	//fmt.Println("x = ", obj.x, ", y = ", obj.y)
	i := (obj.x - obj.startX) / obj.gridW
	j := (obj.y - obj.startY) / obj.gridH
	//fmt.Printf("(%d, %d)\n", i, j)

	if i >= 0 && i <= 7 && j >= 0 && j <= 7 {
		fmt.Printf("(%d, %d)\n", i, j)

		obj.chess[i][j] = obj.currentRole

		//刷新绘图区域（整个窗口）
		obj.window.QueueDraw()

		//改变角色
		obj.ChangeRole()
	}

}

//鼠标移动事件
func MouseMoveEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MouseMoveEvent Chessboard err")
		return
	}

	//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	x, y := int(event.XRoot)-obj.x, int(event.YRoot)-obj.y
	obj.window.Move(x, y) //窗口移动
}

//鼠标移动事件
func PaintEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MouseMoveEvent Chessboard err")
		return
	}

	//获取画家，设置绘图区域
	painter := obj.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(painter)

	//新建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/bg.jpg", obj.w, obj.h, false)

	//黑白棋pixbuf
	blackPixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/black.png", obj.gridW, obj.gridH, false)
	whitePixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/white.png", obj.gridW, obj.gridH, false)

	//画图
	painter.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	//画黑白棋
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				painter.DrawPixbuf(gc, blackPixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				painter.DrawPixbuf(gc, whitePixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}

	//释放资源
	pixbuf.Unref()
	blackPixbuf.Unref()
	whitePixbuf.Unref()
}

//方法：事件、信号处理
func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件
	//"button-press-event"	:鼠标按下时触发
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标移动事件
	//"motion-notify-event"	按住鼠标移动时触发
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)

	//按钮的信号处理
	obj.buttonClose.Clicked(func() {
		//关闭定时器
		//glib.TimeoutRemove(obj.tipTimerId)

		gtk.MainQuit() //关闭窗口
	})

	obj.buttonMin.Clicked(func() {
		obj.window.Iconify() //最小化窗口
	})

	//绘图相关
	//大小改变事件
	//"configure_event"	窗口大小改变时触发
	obj.window.Connect("configure_event", func() {
		//重新刷图
		obj.window.QueueDraw()
	})

	//绘图事件，  "expose-event"
	obj.window.Connect("expose-event", PaintEvent, obj)

}

//函数：提示功能，实现闪烁效果
func ShowTip(obj *Chessboard) {
	if obj.currentRole == Black { //当前黑子下
		//隐藏白子image
		obj.imageWhite.Hide()

		if obj.imageBlack.GetVisible() == true {
			//原来是显示的，需要隐藏
			obj.imageBlack.Hide()
		} else { //原来是隐藏的，需要显示
			obj.imageBlack.Show()
		}

	} else { //当前白棋下
		//隐藏黑子image
		obj.imageBlack.Hide()

		if obj.imageWhite.GetVisible() == true {
			//原来是显示的，需要隐藏
			obj.imageWhite.Hide()
		} else { //原来是隐藏的，需要显示
			obj.imageWhite.Show()
		}

	}

}

//方式：黑白棋属性相关
func (obj *Chessboard) InitChess() {
	//初始化棋盘
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			obj.chess[i][j] = Empty
		}
	}

	//默认有2个黑子和白子
	obj.chess[3][3] = Black
	obj.chess[4][4] = Black
	obj.chess[4][3] = White
	obj.chess[3][4] = White

	//image都先隐藏
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()

	//默认黑子先下
	obj.currentRole = Black
	//obj.currentRole = White

	//启动定时器
	glib.TimeoutAdd(500, func() bool {
		ShowTip(obj)
		return true
	})

	//倒计时定时器
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))

	glib.TimeoutAdd(1000, func() bool {
		obj.timeNum--
		obj.labelTime.SetText(strconv.Itoa(obj.timeNum))

		if obj.timeNum == 0 {
			//时间到，改变角色
			obj.ChangeRole()
		}

		return true
	})

}

func main() {
	//初始化
	gtk.Init(&os.Args)

	var obj Chessboard //创建结构体变量

	obj.CreateWindow() //创建控件，设置控件属性
	obj.HandleSignal() //事件、信号处理
	obj.InitChess()    //黑白棋属性相关

	obj.window.Show() //显示控件

	//主事件循环
	gtk.Main()
}
