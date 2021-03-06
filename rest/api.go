package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bluff/aimbrain"
	"bluff/webexAPI"
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
	ApiKey:    "29354390-b54f-4fe3-ab92-4558ad2114b5",
	ApiSecret: "Y0j9/DRy3R8c+sI4EyEI6fHYajNpAok9SoRFzj4L+hAD0JNcCkcZ25Ab93bWpi4JMACUpRgQRe2FhctEGFWeVQ==",
}

func AuthenticateHandler(writer http.ResponseWriter, request *http.Request) {
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
		//success TODO
		v := webexAPI.FindAndRemove(authRequest.Token)
		b := webexAPI.Bot{Token: v.BotToken}
		b.AddUserToGroup(v.User, v.Room)
		//tell paul's stuff

		writer.WriteHeader(http.StatusOK)
		return
	}

	writer.WriteHeader(http.StatusInternalServerError)
}

func EnrollHandler(writer http.ResponseWriter, request *http.Request) {
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

	sessionResponse, err := ab.GenerateSession("device", 1080, 1920, enrollRequest.UserId, "system")

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
		return
	}
	writer.WriteHeader(http.StatusInternalServerError)
}
