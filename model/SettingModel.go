package model

import (
	"encoding/json"

	"idv/chris/utils"
)

func UnmarshalSetting(data []byte) (Setting, error) {
	var r Setting
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Setting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Setting struct {
	Host    string  `json:"Host"`
	APIKey  string  `json:"APIKey"`
	Prepare Prepare `json:"Prepare"`
	Task    []Task  `json:"Task"`
}

type Prepare struct {
	Action  string                 `json:"Action"`
	Content map[string]interface{} `json:"Content"`
}

type Task struct {
	Action  string                 `json:"Action"`
	Active  bool                   `json:"Active"`
	Content map[string]interface{} `json:"Content"`
}

func NewSetting() Setting {
	file_content, e := utils.FileRead("./asset/setting.json")
	if e != nil {
		panic(e.Error())
	}
	var data Setting
	e = json.Unmarshal(file_content, &data)
	if e != nil {
		panic(e.Error())
	}
	return data
}
