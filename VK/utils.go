package VK

import (
	"multi-messenger-server/database"
	"strconv"
)

func getAccessTokenFromUserId(userId string) (string, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return "", err
	}

	var vkSession database.VKSession
	database.DB.Where("user_id = ?", userIdInt).First(&vkSession)
	return vkSession.AccessToken, nil
}
