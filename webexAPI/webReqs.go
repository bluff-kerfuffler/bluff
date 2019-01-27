package webexAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(err, "unable to build GET request to %v", method)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Authorization", "Bearer "+b.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to execute GET request to %v", method)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("webex get returned: ", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (b Bot) Post(method string, params map[string]interface{}) ([]byte, error) {
	byt, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", apiURL+method, bytes.NewBuffer(byt))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+b.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("webex post returned: ", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
