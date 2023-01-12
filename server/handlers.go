package server

import (
	"multi-messenger-server/VK"
	"multi-messenger-server/auth"
	"multi-messenger-server/tools"
	"net/http"
)

func Handler_vk_auth(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.Write(tools.ErrorJsonBytes(tools.Bad_request))
		return
	}

	authCode := body["code"].(string)
	auth_data, err := VK.GetVKAuthData(authCode)

	if err != nil {
		w.Write(tools.ErrorJsonBytes(tools.Internal_server_error))
		return
	}

	access_token := auth_data["access_token"].(string)
	//expires := int(auth_data["expires_in"].(float64))
	user_id := int(auth_data["user_id"].(float64))

	user, err := VK.GetVkUser(access_token, user_id)

	if err != nil {
		w.Write(tools.ErrorJsonBytes(tools.Internal_server_error))
		return
	}

	w.Write(tools.EncodeJson(user, tools.Internal_server_error))
}

func Handler_login(w http.ResponseWriter, r *http.Request) {

	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.Write(tools.ErrorJsonBytes(tools.Bad_request))
		return
	}
	if val, ok := body["refresh_token"]; ok {
		response := auth.Login("", "", val.(string))
		w.Write(tools.EncodeJson(response, tools.Internal_server_error))
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Login(login.(string), password.(string), "")
		w.Write(tools.EncodeJson(response, tools.Internal_server_error))
		return
	}
	w.Write(tools.ErrorJsonBytes(tools.Bad_request))
}

func Handler_register(w http.ResponseWriter, r *http.Request) {
	body, err := tools.ParseBody(r.Body)
	if err != nil {
		w.Write(tools.ErrorJsonBytes(tools.Bad_request))
		return
	}
	password, okp := body["password"]
	login, okl := body["login"]
	if okl && okp {
		response := auth.Register(login.(string), password.(string))
		w.Write(tools.EncodeJson(response, tools.Internal_server_error))
		return
	}
	w.Write(tools.ErrorJsonBytes(tools.Bad_request))
}
