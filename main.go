package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"bluff/rest"
	"bluff/webexAPI"
)

// get your own token!
// https://developer.webex.com/docs/api/v1/people/get-my-own-details
// create a YAML file called config.yaml in this dir, containing
// token: <yourtoken>

const integrEndpoint = "/integrate"
const domain = "https://kerfuffler.duckdns.org"
const redirURL = domain + integrEndpoint
const accTokURL = "https://api.ciscospark.com/v1/access_token"

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Failed to find a config in current directory.")
	}
	logrus.Info("Using config file at:", viper.ConfigFileUsed())
}

func main() {
	router := mux.NewRouter()
	botRouter := mux.NewRouter()

	router.HandleFunc("/authenticate/", rest.AuthenticateHandler)
	router.HandleFunc("/enroll/", rest.EnrollHandler)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello there"))
		if err != nil {
			logrus.Panic(err)
		}
		w.WriteHeader(200)
	})

	botRouter.HandleFunc(integrEndpoint, func(w http.ResponseWriter, r *http.Request) {
		handleIntegrate(botRouter, w, r)
	})

	go func() {
		log.Fatal(http.ListenAndServeTLS(":8080",
			viper.GetString("certfile"),
			viper.GetString("keyfile"),
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}))(router)))
	}()
	go func() {
		log.Fatal(http.ListenAndServeTLS(":443",
			viper.GetString("certfile"),
			viper.GetString("keyfile"),
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}))(botRouter)))
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

type IntegrationsResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

func handleIntegrate(mux *mux.Router, w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query()["code"]

	v := make(map[string]interface{})
	v["grant_type"] = "authorization_code"
	v["client_id"] = viper.GetString("client_id")
	v["client_secret"] = viper.GetString("client_secret")
	v["code"] = code[0]
	v["redirection_uri"] = redirURL
	byt, _ := json.Marshal(v)
	req, err := http.NewRequest("POST", accTokURL, bytes.NewBuffer(byt))
	if err != nil {
		fmt.Println("failed to post for integration linking", err)
		return
	}

	req.Header.Set("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("failed to post for integration linking", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("unexpected status code for integration linking: ", resp.StatusCode)
		d, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("error is ", d)
		return
	}
	fmt.Println("worked all good")

	defer resp.Body.Close()

	var ir IntegrationsResponse
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&ir)
	if err != nil {
		fmt.Println("failed to decode json for integration linking", err)
		return
	}

	// TODO: Actually keep track of the shitty refresh tokens and add the fucking refresh logic. its a hackathon, yolo
	// NOTE: adding new bots here
	b := webexAPI.Bot{
		Token: ir.AccessToken,
	}

	webhook := webexAPI.Webhook{
		ServePath: b.Token,                // serve on token path so people cant send garbage shit over
		URL:       viper.GetString("url"), // where we set the hook to
	}
	b.AddWebhook(webhook, mux)
	w.Write([]byte("fuck yes"))
}
