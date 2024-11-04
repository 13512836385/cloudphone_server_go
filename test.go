package main

import "cloudphone_server_go/coa"

func main() {
	for i := 0; i <= 7; i++ {
		// 切换角色
		coa.SwitchRole(i)
		// 进入游戏
		coa.EnterGame()
		// 领取奖励+月卡
		coa.ClaimRewards()
		coa.ClaimRewards()
		// 深渊挑战14次
		for j := 0; j < 14; j++ {
			coa.Abyss(j)
		}
		// 返回角色界面
		coa.TapSwitchRole()
	}
}
