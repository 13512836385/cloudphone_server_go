package coa

import (
	"cloudphone_server_go/control"
	struct_package "cloudphone_server_go/struct"
	"cloudphone_server_go/util"
	"time"
)

type Position struct {
	X int
	Y int
}

var server = control.SetControlSettings(
	struct_package.AcepConfig{
		ProductID: util.GetAcepConfig().ProductID,
		PodIDList: util.GetAcepConfig().PodIDList,
		PodID:     util.GetAcepConfig().PodID,
	},
)

func GetRoleCount(role string) int {

	return 0
}

func SwitchRole(role_index int) {
	// 按顺序切换角色,role_index从1开始
	var role_position = [9]Position{
		{X: 250, Y: 200},
		{X: 250, Y: 350},
		{X: 250, Y: 500},
		{X: 250, Y: 650},
		{X: 250, Y: 800},
		{X: 250, Y: 750},
		{X: 250, Y: 600},
		{X: 250, Y: 450},
		{X: 250, Y: 300},
	}

	if role_index <= 4 {
		// 向上滑动到最上为止
		server.Swipe(400, 300, 400, 1000, 1000)

		// 选择角色
		server.Tap(role_position[role_index].X, role_position[role_index].Y)
	} else {
		// 向下滑动到最下为止
		server.Swipe(400, 1000, 400, 300, 1000)
		// 选择角色
		server.Tap(role_position[role_index].X, role_position[role_index].Y)
	}

	time.Sleep(1 * time.Second)

}

func EnterGame() {
	// 点击进入游戏
	server.Tap(1600, 800)
	time.Sleep(10 * time.Second)
}
