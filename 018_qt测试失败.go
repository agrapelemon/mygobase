package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"runtime"
)

func main() {
	////https://www.cnblogs.com/apocelipes/p/9296754.html
	//widgets.NewQApplication(len(os.Args), os.Args)
	//// left sider
	//splitterLeft := widgets.NewQSplitter2(core.Qt__Horizontal, nil)
	//textTop := widgets.NewQTextEdit2("左部文本", splitterLeft)
	//splitterLeft.AddWidget(textTop)
	//
	//// right sider
	//splitterRight := widgets.NewQSplitter2(core.Qt__Vertical, splitterLeft)
	//textRight := widgets.NewQTextEdit2("右部文本", splitterRight)
	//textbuttom := widgets.NewQTextEdit2("下部文本", splitterLeft)
	//splitterRight.AddWidget(textRight)
	//splitterRight.AddWidget(textbuttom)
	//
	//splitterLeft.SetWindowTitle("splitter")
	//splitterLeft.Show()
	//
	//widgets.QApplication_Exec()

	ui.Run(func() {
		//https://studygolang.com/p/goqt
		info := fmt.Sprintf("Hello GoQt Version %v \ngo verison %v %v/%v", ui.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH)
		fmt.Println(info)
		fmt.Println(ui.Version())

		//lable := ui.NewLabel()
		//lable.SetText(info)
		//
		//hbox := ui.NewHBoxLayout()
		//hbox.AddWidget(lable)
		//
		//widget := ui.NewWidget()
		//widget.SetLayout(hbox)
		//widget.Show()
	})
}
