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

type EnrollRequest struct {
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
		//tell paul's stuff
	}
}

func enrollHandler(writer http.ResponseWriter, request *http.Request) {
	d := json.NewDecoder(request.Body)
	var enrollRequest EnrollRequest
	err := d.Decode(&enrollRequest)

	if err != nil {
		fmt.Println("error decoding incoming enrollRequest", err)
		return
	}

	ab := &aimbrain.AimBrain{
		ApiKey:    viper.GetString("aimbrain_api"),
		ApiSecret: viper.GetString("aimbrain_secret"),
	}

	sessionResponse, err := ab.GenerateSession("device", 0, 0, enrollRequest.user_id, "system")

	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	enrollResponse, err := ab.EnrollUser(sessionResponse.Session, enrollRequest.image)

	if err != nil {
		fmt.Println("error enrolling user", err)
		return
	}

	if enrollResponse.ImagesCount > 0 {
		//success
		//tell paul's stuff
	}
}

