package api

import (
	"bytes"
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
===== RESPONSE =====
`
	log_msg = fmt.Sprintf(log_msg, action, body)

	defer func() {
		file_name := strings.ReplaceAll(action, "/", "_")
		utils.FileWriteAppend(filepath.Join("./dist", dir_path), file_name+".txt", []byte(log_msg))
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
	fmt.Printf("\nBeforeJSON:\n%s\n", string(output))

	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)

	fmt.Println("\nResponse:")
	content := buf.String()
	fmt.Println(content)

	log_msg += content

	return data
}
