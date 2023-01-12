package server

import (
	"multi-messenger-server/VK"
	"multi-messenger-server/auth"
	"multi-messenger-server/tools"
	"net/http"
)

func HandlerVkAuth(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.Bad_request)
		return
	}

	authCode := body["code"].(string)
	authData, err := VK.GetVKAuthData(authCode)

	if err != nil {
		w.WriteHeader(tools.Internal_server_error)
		return
	}

	accessToken := authData["access_token"].(string)
	//expires := int(auth_data["expires_in"].(float64))
	userId := int(authData["user_id"].(float64))

	user, err := VK.GetVkUser(accessToken, userId)

	if err != nil {
		w.WriteHeader(tools.Bad_request)
		return
	}
	w.WriteHeader(tools.Ok)
	w.Write(tools.EncodeJson(user))
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {

	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.Bad_request)
		return
	}
	if val, ok := body["refresh_token"]; ok {
		response := auth.Login("", "", val.(string))
		w.WriteHeader(response["status"].(int))
		w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Login(login.(string), password.(string), "")
		w.WriteHeader(response["status"].(int))
		w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
		return
	}
	w.WriteHeader(tools.Bad_request)
}

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(tools.Bad_request)
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Register(login.(string), password.(string))
		w.WriteHeader(response["status"].(int))
		w.Write(tools.EncodeJson(response["data"].(map[string]interface{})))
		return
	}
	w.WriteHeader(tools.Bad_request)
}
