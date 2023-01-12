package tools

import (
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
