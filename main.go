package main

import (
	"fmt"

	"idv/chris/model"
	"idv/chris/utils"
)

func main() {
	setting := model.NewSetting()
	fmt.Println(setting)

	// now := time.Now().Unix()

	fmt.Println(utils.GenSign(setting.Prepare.Content, setting.APIKey))
}
