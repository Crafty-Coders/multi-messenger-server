package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"text/template"
)

func GetTextFromTemplate(tmp string, data map[string]interface{}) (string, error) {
	t := template.Must(template.New("url").Parse(tmp))
	builder := &strings.Builder{}
	err := t.Execute(builder, data)
	if err != nil {
		return "", err
	}
	url := builder.String()
	return url, nil
}

func ParseBody(body io.ReadCloser) (map[string]interface{}, error) {
	parsedBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var bodyMap map[string]interface{}
	json.Unmarshal(parsedBody, &bodyMap)
	return bodyMap, nil
}

func CreateGETQueryFromTemplate(tmp string, data map[string]interface{}) (map[string]interface{}, error) {
	url, err := GetTextFromTemplate(tmp, data)

	if err != nil {
		return nil, err
	}
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ParseBody(resp.Body)
	return body, nil
}
