package api

import (
	"encoding/json"
	"fmt"
	"idv/chris/utils"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	dir_path string = ""
)

func init() {
	dir_path = time.Now().Format("20060102")
}

func ReqTask(host, action, token, body string) any {
	log_msg := `
===== ACTION =====
%s
===== BODY =====
%s
===== =====

`

	defer func() {
		file_name := strings.ReplaceAll(action, "/", "_")
		utils.FileWriteAppend(filepath.Join("./", dir_path), file_name+".txt", []byte(fmt.Sprintf(log_msg, action, body)))
	}()

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
	log_msg += string(output)

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
