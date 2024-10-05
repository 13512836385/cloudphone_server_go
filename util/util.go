package util

import (
	struct_package "cloudphone_server_go/struct"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/volcengine/volc-sdk-golang/base"
	ACEP "github.com/volcengine/volc-sdk-golang/service/acep"
)

func GetConfig() (string, string, error) {
	homeDir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	configPath := filepath.Join(homeDir, ".volc", "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", "", err
	}
	var config struct_package.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", "", err
	}
	return config.AK, config.SK, nil
}

func GetAcepServer() *ACEP.ACEP {
	// 若通过方式二或者方式三方式设置ak、sk会自动读取，不用直接指定设置ak、sk
	service := ACEP.NewInstance()
	// 强烈建议不要把 VOLC_ACCESSKEY 和 VOLC_SECRETKEY 保存到工程代码里，否则可能导致 AccessKey 泄露，威胁您账号下所有资源的安全。
	// 本示例通过从环境变量中读取 VOLC_ACCESSKEY 和 VOLC_SECRETKEY，来实现 API 访问的身份验证。运行代码示例前，请配置环境变量 VOLC_ACCESSKEY 和 VOLC_SECRETKEY
	ak, sk, err := GetConfig()
	if err != nil {
		fmt.Println("Error loading AK and SK:", err)
		return nil
	}
	// fmt.Println("Access Key:", ak)
	// fmt.Println("Secret Key:", sk)

	// fmt.Println(ak)
	// fmt.Println(sk)

	//如果需直接指定ak、sk，则通过以下代码设置
	service.SetCredential(base.Credentials{
		AccessKeyID:     ak,
		SecretAccessKey: sk,
	})
	return service
}
