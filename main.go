package main

import (
	"encoding/json"
	"time"

	"idv/chris/api"
	"idv/chris/model"
	"idv/chris/utils"
)

const (
	OK string = "OK"
)

func toJSON(source any) string {
	content, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func main() {
	setting := model.NewSetting()

	tmp := setting.Prepare.Content
	tmp["Timestamp"] = time.Now().Unix()
	tmp["Sign"] = utils.GenSign(tmp, setting.APIKey)

	auth_res := api.ReqAuth(setting.Host+setting.Prepare.Action, toJSON(tmp))
	status := auth_res["Status"]
	if status != OK {
		panic(status)
	}
	token := auth_res["Data"].(map[string]interface{})["Token"].(string)

	for _, v := range setting.Task {
		tmp := v.Content
		tmp["Timestamp"] = time.Now().Unix()
		tmp["Sign"] = utils.GenSign(tmp, setting.APIKey)

		api.ReqTask(
			setting.Host,
			v.Action,
			token,
			toJSON(tmp),
		)
	}
}
