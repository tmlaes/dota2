package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func topLayout() *fyne.Container {
	loginBtn = widget.NewButton("login", func() {
		login()
	})
	top = container.NewHBox(layout.NewSpacer(), timeLayout(), layout.NewSpacer(), loginBtn)

	return container.NewVBox(top, layout.NewSpacer(), canvas.NewLine(bgc))
}

func timeLayout() *widget.Label {
	timesLay := widget.NewLabelWithData(times)
	go freshTimes()
	return timesLay
}

func middleBox() *fyne.Container {
	return container.NewHBox(leftLayout(), middleLayout(), layout.NewSpacer(), rightLayout())
}

func leftLayout() *fyne.Container {

	startBtn := widget.NewButton("Start\n(F1)", func() {
		start()
	})
	purseBtn := widget.NewButton("Pause\n(F2)", func() {
		pause()
	})
	stopBtn := widget.NewButton("Stop\n(F3)", func() {
		stop()
	})
	return container.NewGridWrap(fyne.Size{100, 188}, startBtn, purseBtn, stopBtn)
}

func middleLayout() *widget.Label {
	return widget.NewLabel("middle")
}

func rightLayout() *fyne.Container {
	richText := widget.NewLabelWithData(logs)
	richText.Wrapping = fyne.TextWrapWord
	return container.NewGridWrap(fyne.NewSize(300, winH), richText)
}
