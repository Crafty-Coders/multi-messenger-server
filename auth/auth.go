package auth

import (
	"golang.org/x/crypto/bcrypt"
	"multi-messenger-server/database"
	"multi-messenger-server/tools"
	"time"
)

func Register(login string, password string) map[string]interface{} {
	var users []database.User

	database.DB.Where("Login = ?", login).Limit(1).Find(&users)

	if len(users) > 0 {
		return map[string]interface{}{
			"status": tools.BadRequest,
			"data": map[string]interface{}{
				"message": "Already exists",
			},
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return map[string]interface{}{
			"status": tools.InternalServerError,
			"data": map[string]interface{}{
				"message": "Err",
			},
		}
	}

	user := database.User{
		Login:    login,
		Password: string(hashedPassword),
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

	database.DB.Where("refresh_token = ?", refreshToken).Limit(1).Find(&sessions)

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
				"user": map[string]interface{}{
					"id": session.UserId,
				},
			},
		}
	}

	return map[string]interface{}{
		"status": tools.NotFound,
		"data": map[string]interface{}{
			"message": "Session not found",
		},
	}
}

func sessionStart(login string, password string) map[string]interface{} {

	var users []database.User

	database.DB.Where("Login = ?", login).Find(&users)

	for _, u := range users {
		if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
			return map[string]interface{}{
				"status": tools.Unauthorized,
				"data": map[string]interface{}{
					"message": "Incorrect password",
				},
			}
		}
		userId := u.Id
		accessToken := generateToken()
		refreshToken := generateToken()
		clearSessions(userId)
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
				"user_id":       userId,
			},
		}
	}

	return map[string]interface{}{
		"status": tools.NotFound,
		"data": map[string]interface{}{
			"message": "User not found",
		},
	}
}
