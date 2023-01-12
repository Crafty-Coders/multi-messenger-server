package main

import (
	"fmt"
	"github.com/rs/cors"
	"multi-messenger-server/config"
	"multi-messenger-server/database"
	"multi-messenger-server/server"
	"net/http"
)

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

	mux.HandleFunc("/vk_auth", server.Handler_vk_auth)

	mux.HandleFunc("/login", server.Handler_login)

	mux.HandleFunc("/register", server.Handler_register)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":5001", handler)
}
