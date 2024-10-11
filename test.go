package main

import "cloudphone_server_go/coa"

func main() {
	for i := 4; i <= 7; i++ {
		// 切换角色
		coa.SwitchRole(i)
		// 进入游戏
		coa.EnterGame()
		// 领取奖励
		coa.ClaimRewards()
		// 深渊挑战14次
		for j := 0; j < 14; j++ {
			coa.Abyss(j)
		}
		// 返回角色界面
		coa.TapSwitchRole()
	}
}

// import (
// 	"fmt"

// 	"gocv.io/x/gocv"
// )

// func main() {
// 	fmt.Printf("gocv version: %s\n", gocv.Version())
// 	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())

// 	window := gocv.NewWindow("Hello")

// 	img := gocv.IMRead("test.jpg", gocv.IMReadColor)

// 	if img.Empty() {
// 		fmt.Printf("Error reading image from: %v\n", "lena.jpg")
// 		return
// 	}

// 	for {
// 		window.IMShow(img)
// 		if window.WaitKey(1) >= 0 {
// 			break
// 		}
// 	}
// }
