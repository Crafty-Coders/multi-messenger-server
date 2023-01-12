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

	mux.HandleFunc("/vk_auth", server.HandlerVkAuth)

	mux.HandleFunc("/login", server.HandlerLogin)

	mux.HandleFunc("/register", server.HandlerRegister)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":5001", handler)
	fmt.Println("Error listening server:\n", err)
}
