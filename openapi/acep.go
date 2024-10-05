package acep_openapi

import (
	"cloudphone_server_go/util"
	"context"
	"encoding/json"
	"fmt"

	ACEP "github.com/volcengine/volc-sdk-golang/service/acep"
)

var service = util.GetAcepServer()

func RunSyncCommand(ProductID string, PodIDList []string, Command string) *ACEP.RunSyncCommandRes {
	body := &ACEP.RunSyncCommandBody{
		// 实例所归属的业务 ID，可在**云手机控制台 > 业务管理 > 业务详情**中获取。
		ProductID: ProductID,
		// 实例 ID，可通过调用 [ListPod](URL_ADDRESS 接口获取。
		PodIDList: PodIDList,
		// 执行的命令，支持的命令详见 [命令列表](URL_ADDRESS
		Command: Command,
	}
	resp, err := service.RunSyncCommand(context.Background(), body)
	if err != nil {
		fmt.Printf("error %v", err)
		panic(err)
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
	}
	return resp
}

func screen_shot(ProductID string, PodID string) {
	body := &ACEP.ScreenShotBody{}
	body.ProductID = ProductID
	body.PodID = PodID
	resp, err := service.ScreenShot(context.Background(), body)
	if err != nil {
		fmt.Printf("error %v", err)
		panic(err)
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
	}
}
