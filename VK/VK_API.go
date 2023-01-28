package VK

import (
	"fmt"
	"multi-messenger-server/tools"
	"net/http"
)

func GetChats(userId string) {

	accessToken, err := getAccessTokenFromUserId(userId)
	if err != nil {
		//
		return
	}

	resp, err := http.Get("https://api.vk.com/method/messages.getConversations?access_token=" + accessToken)
	if err != nil {
		return
	}

	body, err := tools.ParseBody(resp.Body)
	if err != nil {
		//
		return
	}

	fmt.Println(body)
}
