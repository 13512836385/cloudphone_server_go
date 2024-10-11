package coa

func TapSkill(index int) {
	// index为第几个技能
	if index == 0 {
		server.Tap(1500, 970)
	} else if index == 1 {
		server.Tap(1500, 800)
	} else if index == 2 {
		server.Tap(1620, 700)
	} else if index == 3 {
		server.Tap(1800, 700)
	} else if index == 4 {
		server.Tap(1800, 500)
	}
}

func Battle() {
	// 实现简单的战斗系统
	// 每个技能放一轮
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			TapSkill(j)
		}
	}
}
