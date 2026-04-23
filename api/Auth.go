package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ReqAuth(addr, body string) map[string]interface{} {
	cmd_statement := exec.Command(
		"curl", "-s",
		"-X", "POST", addr,
		"-d", body,
	)
	fmt.Println("Auth:\n", cmd_statement.String())

	output, err := cmd_statement.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:\n", string(output))

	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}
	return data
}
