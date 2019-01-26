package webexAPI

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Bot struct {
	Token string
}

const apiURL = "https://api.ciscospark.com/v1/"

var client = &http.Client{
	Transport:     nil,
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       time.Second * 5,
}

func (b Bot) Get(method string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", apiURL+method, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to build get request to %v", method)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Authorization", "Bearer "+b.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to execute get request to %v", method)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (b Bot) Post(method string, params url.Values) ([]byte, error) {
	//b := bytes.Buffer{}
	//w := multipart.NewWriter(&b)
	//part, err := w.CreateFormFile(fileType, filename)
	//if err != nil {
	//	return nil, err
	//}
	//_, err = io.Copy(part, file)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = w.Close()
	//if err != nil {
	//	return nil, err
	//}

	req, err := http.NewRequest("POST", apiURL+method, nil)
	if err != nil {
		logrus.WithError(err).Errorf("failed to send to %v func", method)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Authorization", "Bearer "+b.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
