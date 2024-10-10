package format

import (
	"fmt"
	"strings"
)

// StringToMap 字符串转map
// 例：aa=aa,bb=bb转换成{"aa":"aa", "bb":"bb"}
func StringToMap(str string) map[string]string {
	m := make(map[string]string)
	if str == "" {
		return m
	}

	strS := strings.Split(str, ",")
	for i := 0; i < len(strS); i++ {
		kv := strings.Split(strS[i], "=")
		if _, ok := m[kv[0]]; ok {
			continue
		}

		m[kv[0]] = kv[1]
	}

	return m
}

// MapToString map转字符器
// {"aa":"aa", "bb":"bb"} 转换成 aa=aa,bb=bb
func MapToString(lab map[string]string) string {
	label := make([]string, 0)

	for k, v := range lab {
		label = append(label, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(label, ",")
}
