package auth

import (
	"math/rand"
	"multi-messenger-server/database"
	"strings"
)

func clearSessions(userId uint64) {
	database.DB.Unscoped().Delete(&database.AuthSession{}, "user_id = ?", userId)
}

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
