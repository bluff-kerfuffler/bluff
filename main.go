package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"bluff/webexAPI"
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
	b := webexAPI.Bot{
		Token: token,
	}
	//fmt.Println(b.GetMe())
	//r, _ := b.GetRooms()
	//for _, r := range r.Items {
	//	//fmt.Println(r.Title)
	//	if r.Title == "Bluff" {
	//		msgs, err := b.GetMessages(r.Id, 10)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		for _, m := range msgs.Items {
	//			if len(m.MentionedPeople) > 0 {
	//				fmt.Println(m.MentionedPeople)
	//			}
	//		}
	//	}
	//	//spew.Dump(b.GetUser(r.CreatorId))
	//}

	// start webserver in goroutine
	// start bot in go routine
	// create endless while loop + signal handlers in main loop
}
