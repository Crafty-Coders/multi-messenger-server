package server

import (
	"errors"
	"fmt"
	"multi-messenger-server/auth"
	"multi-messenger-server/tools"
	"net/http"
	"strings"
)

func getDataFromUri(uri string) (string, string, string, error) {
	accessToken := ""
	userId := ""
	expires := ""

	idx := strings.Index(uri, "access_token=") + len("access_token=")
	for i := idx; i < len(uri); i++ {
		if string(uri[i]) == "&" {
			break
		}
		accessToken += string(uri[i])
	}

	idx = strings.Index(uri, "user_id=") + len("user_id=")
	for i := idx; i < len(uri); i++ {
		if string(uri[i]) == "&" {
			break
		}
		userId += string(uri[i])
	}

	idx = strings.Index(uri, "expires_in=") + len("expires_in=")
	for i := idx; i < len(uri); i++ {
		if string(uri[i]) == "&" {
			break
		}
		expires += string(uri[i])
	}

	if accessToken == "" || userId == "" || expires == "" {
		return "", "", "", errors.New("ABOBA")
	}

	return accessToken, userId, expires, nil
}

func HandlerVkAuth(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		return
	}
	userId := body["user_id"].(string)
	uri := body["uri"].(string)
	accessToken, userId, expires, err := getDataFromUri(uri)

	if err != nil {
		return
	}

	fmt.Println(accessToken, userId, expires)
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {

	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.BadRequest)
		return
	}
	if val, ok := body["refresh_token"]; ok {
		response := auth.Login("", "", val.(string))
		w.WriteHeader(response["status"].(int))
		encodedData, err := tools.EncodeJson(response["data"].(map[string]interface{}))
		if err != nil {
			w.WriteHeader(tools.InternalServerError)
			return
		}
		_, err = w.Write(encodedData)

		if err != nil {
			fmt.Println("Error writing response:\n", err)
		}
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Login(login.(string), password.(string), "")
		w.WriteHeader(response["status"].(int))
		encodedData, err := tools.EncodeJson(response["data"].(map[string]interface{}))
		if err != nil {
			w.WriteHeader(tools.InternalServerError)
			return
		}
		_, err = w.Write(encodedData)
		if err != nil {
			fmt.Println("Error writing response:\n", err)
		}
		return
	}
	w.WriteHeader(tools.BadRequest)
}

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.BadRequest)
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Register(login.(string), password.(string))
		w.WriteHeader(response["status"].(int))
		encodedData, err := tools.EncodeJson(response["data"].(map[string]interface{}))
		if err != nil {
			w.WriteHeader(tools.InternalServerError)
			return
		}
		_, err = w.Write(encodedData)
		if err != nil {
			fmt.Println("Error writing response:\n", err)
		}
		return
	}
	w.WriteHeader(tools.BadRequest)
}
