package webexAPI

import (
	"net/http"
	"net/url"
)

const webexIntegrationURL = "https://api.ciscospark.com/v1/authorize"

func (b Bot) GetIntegrationLoginURL() (string, error) {
	v := url.Values{}
	v.Set("response_type", "code") // this is preset
	v.Set("client_id", "") // TODO get this from integration
	v.Set("redirect_uri", "kerfuffler.duckdns.org/integrate")
	v.Set("scope", "spark:all")
	v.Set("state", "helloFriend:)") // TODO figure this bit out
	r, err := http.NewRequest("GET", webexIntegrationURL, nil)
	if err != nil {
		return "", err
	}
	r.URL.RawQuery = v.Encode()
	return r.URL.String(), nil
}
