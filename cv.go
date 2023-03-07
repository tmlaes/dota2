package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

// 在图像中搜索模板，并返回匹配位置
func findTemplate(img gocv.Mat, template gocv.Mat) image.Point {
	// 使用模板匹配算法
	result := gocv.NewMat()
	defer result.Close()
	gocv.MatchTemplate(img, template, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	// 获取匹配结果的最大值和位置
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)

	// 判断匹配结果是否可信
	if maxVal > 0.8 {
		return maxLoc
	} else {
		return image.Point{-1, -1}
	}
}

func cv() {
	// 读取原图和模板图
	img := gocv.IMRead("./imgs/2.png", gocv.IMReadColor)
	template := gocv.IMRead("./imgs/2_.png", gocv.IMReadColor)
	defer img.Close()
	defer template.Close()

	// 在原图中搜索模板
	loc := findTemplate(img, template)
	fmt.Println(loc.X, "====", loc.Y)


	// 绘制匹配结果
	//if loc.X != -1 && loc.Y != -1 {
	//	// 绘制矩形框
	//	size := template.Size()
	//	rect := image.Rect(loc.X, loc.Y, loc.X+size[1], loc.Y+size[0])
	//	gocv.Rectangle(&img, rect, color.RGBA{0, 255, 0, 0}, 2)
	//	// 显示结果
	//	window := gocv.NewWindow("MatchResult")
	//	defer window.Close()
	//	window.IMShow(img)
	//	gocv.WaitKey(0)
	//} else {
	//	fmt.Println("未找到匹配的区域")
	//}
}