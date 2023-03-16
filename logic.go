package main

import (
	"fyne.io/fyne/v2/widget"
	"os"
	"strconv"
)

func login() {
	top.Remove(loginBtn)
	userName.Set("test")
	top.Add(widget.NewLabelWithData(userName))
}

func start() {
	logs.Set("start")
	point := cv()
	logs.Set(strconv.Itoa(point.X) + "====" + strconv.Itoa(point.Y))
}

func pause() {
	dir, errs := os.Getwd()
	if errs != nil {
		logs.Set(errs.Error())
	}
	logs.Set(dir)
}

func stop() {
	logs.Set("stop")
}
