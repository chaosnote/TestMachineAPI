package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ReqTask(host, action, token, body string) any {
	fmt.Printf("\n\n===== Action(%s) =====\n\n", action)
	cmd_statement := exec.Command(
		"curl", "-s",
		"-X", "POST", fmt.Sprintf("%s%s", host, action),
		"-H", fmt.Sprintf(`Token: %s`, token),
		"-d", body,
	)
	fmt.Println(cmd_statement.String())

	output, err := cmd_statement.Output()
	if err != nil {
		panic(err)
	}

	// 顯示符合，但 Unicode 值未處理
	// var pretty_json bytes.Buffer
	// err = json.Indent(&pretty_json, output, "", "  ")

	fmt.Printf("\nBeforeJSON:\n%s\n", string(output))

	var data any
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}

	var content []byte
	content, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nResponse:")
	fmt.Println(string(content))

	return data
}
