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
	return "https://sheltered-bayou-37230.herokuapp.com/verify?token=" + token + "&user_id=" + user
}

func generateToken() string {
	return uuid.New().String()
}
