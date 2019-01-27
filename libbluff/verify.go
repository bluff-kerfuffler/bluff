package libbluff

import "github.com/google/uuid"

type verify struct {
	token   string
	user    string
	channel string
}

func getVerifAddURL(user string, channel string) string {
	token := generateToken()
	storeVerifAdd(verify{token: token, user: user, channel: channel})
	return token
}

func generateToken() string {
	return uuid.New().String()
}
