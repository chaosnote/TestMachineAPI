package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ReqAuth(host, body string) map[string]interface{} {
	cmd_statement := exec.Command(
		"curl", "-s",
		"-X", "POST", host,
		"-d", body,
	)
	fmt.Println("Auth\n", cmd_statement.String())

	output, err := cmd_statement.Output()
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}
	return data
}
