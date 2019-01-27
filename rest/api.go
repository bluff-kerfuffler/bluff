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

func AuthenticateHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	if (*request).Method == "OPTIONS" {
		return
	}

	d := json.NewDecoder(request.Body)
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

		writer.WriteHeader(http.StatusOK)
		return
	}

	writer.WriteHeader(http.StatusInternalServerError)
}

func EnrollHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	if (*request).Method == "OPTIONS" {
		return
	}


	d := json.NewDecoder(request.Body)
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

	fmt.Println("HERE")

	if enrollResponse.ImagesCount > 0 {
		//success
		//TODO tell paul's stuff
		//tell paul's stuff

		writer.WriteHeader(http.StatusOK)
	}
	writer.WriteHeader(http.StatusInternalServerError)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}