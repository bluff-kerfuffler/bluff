package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	token = viper.GetString("token")


	// start webserver in goroutine
	// start bot in go routine
	// create endless while loop + signal handlers in main loop
}
