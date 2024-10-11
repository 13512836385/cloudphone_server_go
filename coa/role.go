package coa

import "time"

// 一键领取奖励
func ClaimRewards() {
	// 点击领取
	server.Tap(950, 850)
	// 点击空白位置
	server.Tap(950, 850)
}

// 点击战斗模块
func TapBattle() {
	server.Tap(1825, 430)
}

// 点击冒险玩法
func TapAdventure() {
	server.Tap(400, 400)
}

func WaitLoading() {
	// 等待读屏加载地图
	time.Sleep(10 * time.Second)
}
func TapReturn() {
	// 点击返回
	server.Tap(100, 50)
}
func TapCurrentMap() {
	// 点击地图
	server.Tap(1800, 130)
}
func TapWorldMap() {
	// 点击当前地图中的世界地图
	server.Tap(1600, 925)
}

func TapMapByName(name string) {
	// 根据名称点击地图，暂时只支持鲁米村
	if name == "鲁米村" {
		server.Tap(1100, 850)
	}
}

func TapEnsure() {
	// 点击确定
	server.Tap(1250, 720)
}

func TapExitReward() {
	// 点击退出副本奖励界面
	server.Tap(1000, 200)
}

func TapExitInstance() {
	// 点击退出副本
	server.Tap(1080, 60)
	// 确定
	TapEnsure()
	// 等待读屏加载地图
	WaitLoading()
}
func TapMore() {
	// 点击更多
	server.Tap(1830, 550)
}
func TapSettings() {
	// 点击设置
	TapMore()
	server.Tap(1860, 900)
}
func TapSwitchRole() {
	// 点击切换角色
	TapSettings()
	server.Tap(1770, 940)
}

// 进行深渊挑战
func Abyss(index int) {
	// index参数为深渊挑战的执行次数

	// 首次挑战
	if index == 0 {
		// 重置角色位置保证当前角色不在虚空表域
		// 1.点击地图
		TapCurrentMap()
		// 2.点击世界地图
		TapWorldMap()
		// 3.选择鲁米村
		TapMapByName("鲁米村")
		// 4.点击返回地图,保证当前角色已经在鲁米村时程序正常执行
		TapReturn()
		// 等待地图加载
		WaitLoading()
		// 点击战斗玩法
		TapBattle()
		// 点击冒险玩法
		TapAdventure()
		// 选择深渊暗域
		server.Tap(170, 870)
		// 选择双倍or正常
		server.Tap(900, 1000)
		// 点击挑战传送到深渊地图
		server.Tap(1700, 1000)
		// 等待读屏加载地图
		WaitLoading()
		// 再次选择挑战等待自动调整方向
		TapBattle()
		TapAdventure()
		server.Tap(1700, 1000)
		time.Sleep(3 * time.Second)
		// 移动到npc位置
		server.Move(290, 400, 5000)
		// time.Sleep(10 * time.Second)
		// 点击确定挑战
		TapEnsure()
	} else {
		// 点击战斗玩法
		TapBattle()
		// 点击冒险玩法
		TapAdventure()
		// 深渊挑战
		server.Tap(1700, 1000)
	}
	// 等待读屏加载地图
	WaitLoading()
	// 移动触发战斗
	server.Move(290, 400, 3500)
	// time.Sleep(100 * time.Second)
	// 战斗系统
	Battle()
	// 退出副本奖励界面
	TapExitReward()
	// 再点击一次防止由于战斗操作点击选中奖励物品介绍
	TapExitReward()
	// 再点击一次防止卡在柱子界面
	time.Sleep(10 * time.Second)
	TapExitReward()
	// 离开副本
	TapExitInstance()
}
