package main

import (
	"encoding/json"
	"time"

	"idv/chris/api"
	"idv/chris/model"
	"idv/chris/utils"
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

	api.ReqAuth(setting.Host, toJSON(tmp))
}
