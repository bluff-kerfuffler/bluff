package rest

import (
	"bluff/aimbrain"
	"bluff/libbluff"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type AuthRequest struct {
	token string
	user_id string
	image string
}

func InitRestAPI() {
	http.HandleFunc("/authenticate/", authenticateHandler)

}

func authenticateHandler(writer http.ResponseWriter, request *http.Request) {
	d := json.NewDecoder(request.Body)
	var authRequest AuthRequest
	err := d.Decode(&authRequest)

	if err != nil {
		fmt.Println("error decoding incoming authRequest", err)
		return
	}

	ab := &aimbrain.AimBrain{
		ApiKey:    viper.GetString("aimbrain_api"),
		ApiSecret: viper.GetString("aimbrain_secret"),
	}

	sessionResponse, err := ab.GenerateSession("device", 0, 0, authRequest.user_id, "system")

	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	authResponse, err := ab.Authuser(sessionResponse.Session, authRequest.image)

	if err != nil {
		fmt.Println("error authenticating user", err)
		return
	}

	if authResponse.Score > 0.2 {
		//success
		libbluff.FindAndRemove(authRequest.token)

	}
}

