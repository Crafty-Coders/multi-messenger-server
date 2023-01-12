package tools

import "encoding/json"

func ErrorJsonBytes(errorCode int) []byte {
	errBytes, _ := json.Marshal(map[string]interface{}{
		"status":  errorCode,
		"message": "",
	})
	return errBytes
}

func EncodeJson(data map[string]interface{}, defaultErrorCode int) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return ErrorJsonBytes(defaultErrorCode)
	}
	return b
}
