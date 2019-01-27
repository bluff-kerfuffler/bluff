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

	ab := &aimbrain.AimBrain{
		ApiKey:    viper.GetString("aimbrain_api"),
		ApiSecret: viper.GetString("aimbrain_secret"),
	}

	sessionResponse, err := ab.GenerateSession("device", 1080, 1920, authRequest.user_id, "system")

	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	authResponse, err := ab.AuthUser(sessionResponse.Session, authRequest.image)

	if err != nil {
		fmt.Println("error authenticating user", err)
		return
	}

	if authResponse.Score > 0.2 {
		//success
		libbluff.FindAndRemove(authRequest.token)
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

	ab := &aimbrain.AimBrain{
		ApiKey:    viper.GetString("aimbrain_api"),
		ApiSecret: viper.GetString("aimbrain_secret"),
	}

	sessionResponse, err := ab.GenerateSession("device", 1080, 1920, enrollRequest.user_id, "system")

	if err != nil {
		fmt.Println("error generating session", err)
		return
	}

	enrollResponse, err := ab.EnrollUser(sessionResponse.Session, enrollRequest.image)

	if err != nil {
		fmt.Println("error enrolling user", err)
		return
	}

	fmt.Println("HERE")

	if enrollResponse.ImagesCount > 0 {
		//success
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