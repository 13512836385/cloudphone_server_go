package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	ACEP "github.com/volcengine/volc-sdk-golang/service/acep"
)

func main() {
	// 若通过方式二或者方式三方式设置ak、sk会自动读取，不用直接指定设置ak、sk
	service := ACEP.NewInstance()

	// 强烈建议不要把 VOLC_ACCESSKEY 和 VOLC_SECRETKEY 保存到工程代码里，否则可能导致 AccessKey 泄露，威胁您账号下所有资源的安全。
	// 本示例通过从环境变量中读取 VOLC_ACCESSKEY 和 VOLC_SECRETKEY，来实现 API 访问的身份验证。运行代码示例前，请配置环境变量 VOLC_ACCESSKEY 和 VOLC_SECRETKEY
	ak, sk, err := loadAKAndSK()
	if err != nil {
		fmt.Println("Error loading AK and SK:", err)
		return
	}
	fmt.Println("Access Key:", ak)
	fmt.Println("Secret Key:", sk)

	fmt.Println(ak)
	fmt.Println(sk)

	//如果需直接指定ak、sk，则通过以下代码设置
	service.SetCredential(base.Credentials{
		AccessKeyID:     ak,
		SecretAccessKey: sk,
	})

	query := &ACEP.DetailPodQuery{
		// 实例所归属的业务 ID，可在**云手机控制台 > 业务管理 > 业务详情**中获取。
		ProductID: `1608456935646957568`,
		// 实例 ID，可通过调用 [ListPod](https://www.volcengine.com/docs/6394/1221468) 接口获取。
		PodID: `7398881170744335116`,
	}

	resp, err := service.DetailPod(context.Background(), query)
	fmt.Printf("product_id %v", resp.Result.DisplayLayoutID)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
