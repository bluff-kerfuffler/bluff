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
	//url := viper.GetString("url")
	//b := webexAPI.Bot{
	//	Token: token,
	//}
	//
	//webhook := webexAPI.Webhook{
	//	Serve:     "0.0.0.0", // localhost so we can grab caddy later
	//	ServePort: 443,       // HTTPS ftw
	//	ServePath: b.Token,   // serve on token path so people cant send garbage shit over
	//	URL:       url,       // where we set the hook to
	//}
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
	//	ApiKey:    "29354390-b54f-4fe3-ab92-4558ad2114b5",
	//	ApiSecret: "Y0j9/DRy3R8c+sI4EyEI6fHYajNpAok9SoRFzj4L+hAD0JNcCkcZ25Ab93bWpi4JMACUpRgQRe2FhctEGFWeVQ==",
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
	//_, err = ab.Authuser(sess.Session, imgBase64Str)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
