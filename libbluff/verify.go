package libbluff

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type verify struct {
	token string
	user  string
	team  string
}

func GetVerifAddURL(user string, team string) string {
	token := generateToken()
	storeVerifAdd(verify{token: token, user: user, team: team})
	r, _ := http.NewRequest("GET", "https://sheltered-bayou-37230.herokuapp.com/verify", nil)
	r.URL.RawQuery = url.Values{"user_id": []string{user}, "token": []string{token}}.Encode()
	return r.URL.String()
}

func generateToken() string {
	return uuid.New().String()
}
