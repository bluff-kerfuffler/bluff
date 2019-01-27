package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// get your own token!
// https://developer.webex.com/docs/api/v1/people/get-my-own-details
// create a YAML file called config.yaml in this dir, containing
// token: <yourtoken>
//var token = ""
//
//func init() {
//	viper.SetConfigName("config")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath(".")
//	if err := viper.ReadInConfig(); err != nil {
//		logrus.Fatal("Failed to find a config in current directory.")
//	}
//	logrus.Info("Using config file at:", viper.ConfigFileUsed())
//}


func main() {
	//token = viper.GetString("token")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_ , err := writer.Write([]byte("Callum's mum's a big fat slag"))
		if err != nil {
			logrus.Panic(err)
		}
		writer.WriteHeader(200)
	})
	logrus.Fatal(http.ListenAndServe(":8080", nil))

	// start webserver in goroutine
	// start bot in go routine
	// create endless while loop + signal handlers in main loop
}
