package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func ReqTask(host, action, token, body string) any {
	fmt.Printf("\n===== Action(%s) =====\n", action)
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

	var pretty_json bytes.Buffer
	err = json.Indent(&pretty_json, output, "", "  ")

	fmt.Println("\nResponse:\n", pretty_json.String())

	var data any
	// err = json.Unmarshal(output, &data)
	// if err != nil {
	// 	panic(err)
	// }

	return data
}
