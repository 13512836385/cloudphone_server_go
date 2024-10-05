package control

import (
	"fmt"

	acep_openapi "cloudphone_server_go/openapi"
)

func Tap(ProductID string, PodIDList []string, X int, Y int) {
	var Command = "input tap " + fmt.Sprint(X) + " " + string(Y)
	acep_openapi.RunSyncCommand(ProductID, PodIDList, Command)
}
