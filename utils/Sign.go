package utils

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func GenSign(source map[string]interface{}, api_key string) string {
	// 移除不需要簽名的欄位
	delete(source, "sign")
	delete(source, "Sign")

	// 取得所有 Key 並進行排序 (ksort)
	keys := make([]string, 0, len(source))
	for k := range source {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var output []string
	for _, k := range keys {
		v := source[k]
		if v == nil || v == "" || v == 0 {
			continue
		}
		output = append(output, fmt.Sprintf("%s=%v", k, v))
	}

	output = append(output, fmt.Sprintf("sKey=%s", api_key))
	tmp := strings.Join(output, "&")
	return fmt.Sprintf("%x", md5.Sum([]byte(tmp)))
}
