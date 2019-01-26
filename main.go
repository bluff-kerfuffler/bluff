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
	// get token
	//token = viper.GetString("token")
	//b := webexAPI.Bot{
	//	Token: token,
	//}
	
	// NOTE: getting bot api
	//fmt.Println(b.GetMe())

	// NOTE: checks room, gets recent messages, sends relevant messages based on them
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
	//				fmt.Println(r.Id)
	//				fmt.Println(m.MentionedPeople[0])
	//
	//			}
	//		}
	//	}
	//	//spew.Dump(b.GetUser(r.CreatorId))
	//}

	// NOTE: sending to a room, mentioning a user and abusing markdown
	//bluffRoom := "Y2lzY29zcGFyazovL3VzL1JPT00vZmM2MTdlZTAtMjE3Yy0xMWU5LTkzMWMtYjE5NzVkYmUwMjdj"
	//michaelID := "Y2lzY29zcGFyazovL3VzL1BFT1BMRS8yY2Q0MDQzOC1mMGE1LTRlMzEtYWM5NS02ODY2YjMzOWQ0Nzg"
	//m, err := b.SendMessageMarkdown(bluffRoom, "we mentioned "+webexAPI.MentionIdMarkdown(michaelID, "this guy.")+ "google.com is a url, **this is bold** and [this](http://example.com) was a hyperlink. ```code is pretty lit```")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}


	// start webserver in goroutine
	// start bot in go routine
	// create endless while loop + signal handlers in main loop
}
