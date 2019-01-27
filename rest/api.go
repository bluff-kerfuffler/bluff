package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"bluff/aimbrain"
	"bluff/libbluff"
)

type AuthRequest struct {
	Token  string
	UserId string
	Image  string
}

type EnrollRequest struct {
	UserId string
	Image  string
}

var ab = &aimbrain.AimBrain{
	ApiKey:    viper.GetString("aimbrain_api"),
	ApiSecret: viper.GetString("aimbrain_secret"),
}

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var authRequest AuthRequest
	err := d.Decode(&authRequest)
	if err != nil {
		fmt.Println("error decoding incoming authRequest", err)
		return
	}

	sessionResponse, err := ab.GenerateSession("device", 1080, 1920, authRequest.UserId, "system")

	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	authResponse, err := ab.AuthUser(sessionResponse.Session, authRequest.Image)

	if err != nil {
		fmt.Println("error authenticating user", err)
		return
	}

	if authResponse.Score > 0.2 {
		//success
		libbluff.FindAndRemove(authRequest.Token)
		//tell paul's stuff
	}
}

func enrollHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var enrollRequest EnrollRequest
	err := d.Decode(&enrollRequest)
	if err != nil {
		fmt.Println("error decoding incoming enrollRequest", err)
		return
	}

	sessionResponse, err := ab.GenerateSession("device", 0, 0, enrollRequest.UserId, "system")
	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	enrollResponse, err := ab.EnrollUser(sessionResponse.Session, enrollRequest.Image)
	if err != nil {
		fmt.Println("error enrolling user", err)
		return
	}

	if enrollResponse.ImagesCount > 0 {
		//success
		//TODO tell paul's stuff
	}
}
