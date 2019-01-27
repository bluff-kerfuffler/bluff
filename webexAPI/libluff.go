package webexAPI

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type Verify struct {
	BotToken string
	Token    string
	User     string
	Room     string
}


var m map[string]Verify

func storeVerifAdd(v Verify) {
	if m == nil {
		m = make(map[string]Verify)
	}
	m[v.Token] = v
}

func FindAndRemove(token string) Verify {
	result := m[token]
	delete(m, token)
	return result
}

func (b Bot) GetVerifAddURL(user string, room string) string {
	token := generateToken()
	storeVerifAdd(Verify{
		BotToken: b.Token,
		Token:    token,
		User:     user,
		Room:     room,
	})
	r, _ := http.NewRequest("GET", "https://sheltered-bayou-37230.herokuapp.com/verify", nil)
	r.URL.RawQuery = url.Values{"user_id": []string{user}, "token": []string{token}}.Encode()
	return r.URL.String()
}

func generateToken() string {
	return uuid.New().String()
}
