package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"

	"bluff/webexAPI"
	"net/http"
)

// get your own token!
// https://developer.webex.com/docs/api/v1/people/get-my-own-details
// create a YAML file called config.yaml in this dir, containing
// token: <yourtoken>
var token = ""

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
	token = viper.GetString("token")
	url := viper.GetString("url")
	b := webexAPI.Bot{
		Token: token,
	}

	webhook := webexAPI.Webhook{
		Serve:     "0.0.0.0", // localhost so we can grab caddy later
		ServePort: 443,       // HTTPS ftw
		ServePath: b.Token,   // serve on token path so people cant send garbage shit over
		URL:       url,		  // where we set the hook to
	}
	// starts webserver in goroutine
	b.StartWebhook(webhook)
	// set webhook URL (happens after server start to avoid missing messages)
	_, err := b.SetFirehoseWebhook("allTheMessages", webhook)
	if err != nil {
		log.Fatal("Failed to start bot due to: ", err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_ , err := writer.Write([]byte("Callum's mum's a big fat slag"))
		if err != nil {
			logrus.Panic(err)
		}
		writer.WriteHeader(200)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))


}
