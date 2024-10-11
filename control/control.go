package control

import (
	"fmt"
	"time"

	acep_openapi "cloudphone_server_go/openapi"
	struct_package "cloudphone_server_go/struct"
)

type Control struct {
	*struct_package.AcepConfig
}

func SetControlSettings(acep_config struct_package.AcepConfig) *Control {
	if acep_config.ProductID == "" {
		panic("product_id is empty")
	}
	if acep_config.PodIDList == nil && acep_config.PodID == "" {
		panic("pod_id_list or pod_id is empty")
	}
	control := &Control{
		AcepConfig: &acep_config,
	}
	return control
}

func (control *Control) Tap(X int, Y int) {
	// 点击屏幕
	var Command = "input tap " + fmt.Sprint(X) + " " + fmt.Sprint(Y)
	acep_openapi.RunSyncCommand(control.ProductID, control.PodIDList, Command)
	time.Sleep(1 * time.Second)
}

func (control *Control) Swipe(X1 int, Y1 int, X2 int, Y2 int, duration int) {
	// 上下滑动，并保持指定秒数
	var Command = "input swipe " + fmt.Sprint(X1) + " " + fmt.Sprint(Y1) + " " + fmt.Sprint(X2) + " " + fmt.Sprint(Y2) + " " + fmt.Sprint(duration)
	acep_openapi.RunSyncCommand(control.ProductID, control.PodIDList, Command)

}

func (control *Control) Move(X int, Y int, duration int) {
	// 从轮盘中心向指定位置移动
	x := 290
	y := 850

	control.Swipe(x, y, X, Y, duration)

}
