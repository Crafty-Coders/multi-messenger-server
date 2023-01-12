package auth

import (
	"math/rand"
	"multi-messenger-server/database"
	"multi-messenger-server/tools"
	"strings"
	"time"
)

func generateToken() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet += strings.ToUpper(alphabet)
	alphabet += "0123456789"
	token := ""
	for i := 0; i < 40; i++ {
		token += string([]rune(alphabet)[rand.Intn(len(alphabet))])
	}
	return token
}

func Register(login string, password string) map[string]interface{} {
	var users []database.User

	database.DB.Where("Login = ?", login).Limit(1).Find(&users)

	if len(users) > 0 {
		return map[string]interface{}{
			"status": tools.Bad_request,
			"data": map[string]interface{}{
				"message": "Already exists",
			},
		}
	}

	user := database.User{
		Login:    login,
		Password: password,
	}

	database.DB.Create(&user)

	return sessionStart(login, password)
}

func Login(login string, password string, refreshToken string) map[string]interface{} {
	if password != "" && login != "" {
		return sessionStart(login, password)
	}
	if refreshToken != "" {
		return sessionRefresh(refreshToken)
	}
	return map[string]interface{}{
		"status": tools.Unauthorized,
		"data": map[string]interface{}{
			"message": "Invalid credentials",
		},
	}
}

func sessionRefresh(refreshToken string) map[string]interface{} {

	var sessions []database.AuthSession

	database.DB.Where("RefreshToken = ?", refreshToken).Limit(1).Find(&sessions)

	if len(sessions) > 0 {
		session := sessions[0]
		accessToken := generateToken()
		newRefreshToken := generateToken()
		database.DB.Model(&session).Updates(database.AuthSession{
			RefreshToken: newRefreshToken,
			AccessToken:  accessToken,
			SessionStart: time.Now(),
		})

		return map[string]interface{}{
			"status": tools.Ok,
			"data": map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": newRefreshToken,
			},
		}
	}

	return map[string]interface{}{
		"status": tools.Not_found,
		"data": map[string]interface{}{
			"message": "Session not found",
		},
	}
}

func sessionStart(login string, password string) map[string]interface{} {

	var users []database.User

	database.DB.Where("Login = ? AND Password = ?", login, password).Find(&users)

	for _, u := range users {
		userId := u.Id
		accessToken := generateToken()
		refreshToken := generateToken()
		session := database.AuthSession{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserId:       userId,
			SessionStart: time.Now(),
		}
		database.DB.Create(&session)

		return map[string]interface{}{
			"status": tools.Ok,
			"data": map[string]interface{}{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		}
	}

	return map[string]interface{}{
		"status": tools.Not_found,
		"data": map[string]interface{}{
			"message": "User not found",
		},
	}
}
