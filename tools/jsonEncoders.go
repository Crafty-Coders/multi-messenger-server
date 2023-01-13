package tools

import "encoding/json"

func EncodeJson(data map[string]interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
