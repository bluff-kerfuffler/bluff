package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

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
	// get token

	//// starts webserver in goroutine
	//b.StartWebhook(webhook)
	//// set webhook URL (happens after server start to avoid missing messages)
	//_, err := b.SetFirehoseWebhook("allTheMessages", webhook)
	//if err != nil {
	//	log.Fatal("Failed to start bot due to: ", err)
	//}
	//
	//// endless main loooooop
	//for {
	//	time.Sleep(1 * time.Second)
	//}

	//ab := &aimbrain.AimBrain{
	//	ApiKey:    viper.GetString("aimbrain_api"),
	//	ApiSecret: viper.GetString("aimbrain_secret"),
	//}
	//sess, err := ab.GenerateSession("anoos", 640, 480, "benny", "tool")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//f, err := os.Open("michael.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//// create a new buffer base on file size
	//fInfo, err := f.Stat()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var size int64 = fInfo.Size()
	//buf := make([]byte, size)
	//
	//fReader := bufio.NewReader(f)
	//fReader.Read(buf)
	//
	//imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	//
	//_, err = ab.EnrollUser(sess.Session, imgBase64Str)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = ab.AuthUser(sess.Session, imgBase64Str)
	//if err != nil {
	//	log.Fatal(err)
	//}

	rest.InitRestAPI()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Callum's mum's a big fat slag"))
		if err != nil {
			logrus.Panic(err)
		}
		w.WriteHeader(200)
	})

	mux.HandleFunc(integrEndpoint, func(w http.ResponseWriter, r *http.Request) {
		handleIntegrate(mux, r)
	})

	log.Fatal(http.Server{
		Addr:    ":8080", // todo check this
		Handler: mux,
	}.ListenAndServe())
}

type IntegrationsResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

func handleIntegrate(mux *http.ServeMux, r *http.Request) {
	code := r.URL.Query()["code"]

	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("client_id", viper.GetString("client_id"))
	v.Set("client_secret", viper.GetString("client_secret"))
	v.Set("code", code[0])
	v.Set("redirection_uri", redirURL)
	byt, _ := json.Marshal(v)
	req, err := http.NewRequest("POST", accTokURL, bytes.NewBuffer(byt))
	if err != nil {
		fmt.Println("failed to post for integration linking", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("failed to post for integration linking", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("unexpected status code for integration linking: ", resp.StatusCode)
		return
	}
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
}
