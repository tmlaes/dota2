package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"gui/theme"
)

var (
	loginBtn              *widget.Button
	top                   *fyne.Container
	times, userName, logs binding.String
)

func init() {
	//读取目录

	//用户名
	userName = binding.NewString()
	//时间
	times = binding.NewString()
	//logs
	logs = binding.NewString()
}

func createWin() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	w := myApp.NewWindow("Hello")
	w.SetMaster()
	w.SetPadded(false)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(winW, winH))
	content := container.NewVBox(topLayout(), middleBox())

	w.SetContent(content)
	w.ShowAndRun()
}
