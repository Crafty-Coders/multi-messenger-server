package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"multi-messenger-server/VK"
	"multi-messenger-server/auth"
	"multi-messenger-server/config"
	"multi-messenger-server/database"
	"multi-messenger-server/tools"
	"net/http"
)

func errorJsonBytes(errorCode int) []byte {
	errBytes, _ := json.Marshal(map[string]interface{}{
		"status":  errorCode,
		"message": "",
	})
	return errBytes
}

func encodeJson(data map[string]interface{}, defaultErrorCode int) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return errorJsonBytes(defaultErrorCode)
	}
	return b
}

func main() {

	if err := config.ReadConfig(); err != nil {
		fmt.Println("Error reading config!!!")
		return
	}

	if err := database.ConnectDB(); err != nil {
		fmt.Println("Error connecting database!!!")
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/vk_auth", func(w http.ResponseWriter, r *http.Request) {
		body, err := tools.ParseBody(r.Body)
		if err != nil {
			w.Write(errorJsonBytes(tools.Bad_request))
			return
		}

		authCode := body["code"].(string)
		auth_data, err := VK.GetVKAuthData(authCode)

		if err != nil {
			w.Write(errorJsonBytes(tools.Internal_server_error))
			return
		}

		access_token := auth_data["access_token"].(string)
		//expires := int(auth_data["expires_in"].(float64))
		user_id := int(auth_data["user_id"].(float64))

		user, err := VK.GetVkUser(access_token, user_id)

		if err != nil {
			w.Write(errorJsonBytes(tools.Internal_server_error))
			return
		}

		w.Write(encodeJson(user, tools.Internal_server_error))
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		body, err := tools.ParseBody(r.Body)
		if err != nil {
			w.Write(errorJsonBytes(tools.Bad_request))
			return
		}
		if val, ok := body["refresh_token"]; ok {
			response := auth.Login("", "", val.(string))
			w.Write(encodeJson(response, tools.Internal_server_error))
			return
		}
		password, okp := body["password"]
		login, okl := body["login"]
		if okl && okp {
			response := auth.Login(login.(string), password.(string), "")
			w.Write(encodeJson(response, tools.Internal_server_error))
			return
		}
		w.Write(errorJsonBytes(tools.Bad_request))
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		body, err := tools.ParseBody(r.Body)
		if err != nil {
			w.Write(errorJsonBytes(tools.Bad_request))
			return
		}
		password, okp := body["password"]
		login, okl := body["login"]
		if okl && okp {
			response := auth.Register(login.(string), password.(string))
			w.Write(encodeJson(response, tools.Internal_server_error))
			return
		}
		w.Write(errorJsonBytes(tools.Bad_request))
	})

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":5001", handler)
}
