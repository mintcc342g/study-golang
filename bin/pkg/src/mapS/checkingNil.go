package maps

import (
	"fmt"
)

// IsNil ... map의 nil을 검사하는 함수
func IsNil(m map[string]interface{}) bool {
	if len(m) == 0 {
		return true
	}

	for k, v := range m {
		val, isKeyExist := m[k]
		if !isKeyExist {
			return true
		}
		if val == nil {
			return true
		}

		switch t := v.(type) {
		case []interface{}:
			if len(v.([]interface{})) == 0 {
				return true
			}
		case []string:
			if len(v.([]string)) == 0 {
				return true
			}
		case string:
			if v.(string) == "" {
				return true
			}
			// add cases...
		default:
			fmt.Printf("Unknown Type [%v]", t)
		}
	}
	return false
}
