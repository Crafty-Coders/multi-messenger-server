package server

import (
	"fmt"
	"multi-messenger-server/VK"
	"multi-messenger-server/auth"
	"multi-messenger-server/tools"
	"net/http"
)

func HandlerVkAuth(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.BadRequest)
		return
	}

	authCode := body["code"].(string)
	authData, err := VK.GetVKAuthData(authCode)

	if err != nil {
		w.WriteHeader(tools.InternalServerError)
		return
	}

	accessToken := authData["access_token"].(string)
	//expires := int(auth_data["expires_in"].(float64))
	userId := int(authData["user_id"].(float64))

	user, err := VK.GetVkUser(accessToken, userId)

	if err != nil {
		w.WriteHeader(tools.BadRequest)
		return
	}
	w.WriteHeader(tools.Ok)
	_, err = w.Write(tools.EncodeJson(user))
	if err != nil {
		fmt.Println("Error writing response:\n", err)
	}
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
		_, err = w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
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
		_, err = w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
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
		_, err := w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
		if err != nil {
			fmt.Println("Error writing response:\n", err)
		}
		return
	}
	w.WriteHeader(tools.BadRequest)
}
