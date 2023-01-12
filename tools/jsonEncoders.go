package tools

import "encoding/json"

func ErrorJsonBytes(errorCode int) []byte {
	errBytes, _ := json.Marshal(map[string]interface{}{
		"status":  errorCode,
		"message": "",
	})
	return errBytes
}

func EncodeJson(data map[string]interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
