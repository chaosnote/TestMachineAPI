package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ReqTask(host, action, token, body string) map[string]interface{} {
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

	fmt.Println("\nResponse:\n", string(output))

	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}
	return data
}
