package tools

import "encoding/json"

func EncodeJson(data map[string]interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
