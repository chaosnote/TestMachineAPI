package utils

import "strings"

func GetValueByPath(source map[string]interface{}, req_path string) any {
	if req_path == "" {
		return source
	}

	// 拆分路徑，例如 "AAA.BBB.CCC" -> ["AAA", "BBB", "CCC"]
	parts := strings.Split(req_path, ".")

	var current interface{} = source

	for _, part := range parts {
		// 檢查當前層級是否為 map
		if m, ok := current.(map[string]interface{}); ok {
			// 檢查 key 是否存在
			if val, exists := m[part]; exists {
				current = val
			} else {
				return nil
			}
		} else {
			// 如果當前層級不是 map，但還有路徑要走，表示路徑不合法
			return nil
		}
	}

	return current
}
