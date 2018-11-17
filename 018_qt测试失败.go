package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"runtime"
)

func main() {
	ui.Run(func() {
		info := fmt.Sprintf("Hello GoQt Version %v \ngo verison %v %v/%v", ui.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH)

		lable := ui.NewLabel()
		lable.SetText(info)

		hbox := ui.NewHBoxLayout()
		hbox.AddWidget(lable)

		widget := ui.NewWidget()
		widget.SetLayout(hbox)
		widget.Show()
	})
}
