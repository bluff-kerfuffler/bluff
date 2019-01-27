package aimbrain

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const aimbrainAPI = "https://api.aimbrain.com:443"

type AimBrain struct {
	ApiKey    string
	ApiSecret string
}

type SessionResponse struct {
	Session   string
	Face      int
	Voice     int
	Behaviour int
}

func (ab *AimBrain) GenerateSession(device string, height, width int, userid, system string) (*SessionResponse, error) {
	//'{"device":<device>, "screenHeight":<height>, "screenWidth":<width>, "userId":<userId>, "system":<system>}'
	v := map[string]interface{}{
		"device":       device,
		"screenHeight": height,
		"screenWidth":  width,
		"userId":       userid,
		"system":       system,
	}
	res, err := ab.actualPostRequest("/v1/sessions", v)
	if err != nil {
		return nil, err
	}

	var sr SessionResponse
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &sr, d.Decode(&sr)
}

type AuthResponse struct {
	Score      int
	Liveliness int
}

func (ab *AimBrain) Authuser(session string, images ...string) (*AuthResponse, error) {
	v := map[string]interface{}{
		"session": session,
		"faces":   images,
	}
	res, err := ab.actualPostRequest("/v1/face/auth", v)
	if err != nil {
		return nil, err
	}

	var ar AuthResponse
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &ar, d.Decode(&ar)
}

type EnrollResponse struct {
	ImagesCount int
}

func (ab *AimBrain) EnrollUser(session string, images ...string) (*EnrollResponse, error) {
	v := map[string]interface{}{
		"session": session,
		"faces":   images,
	}
	res, err := ab.actualPostRequest("/v1/face/enroll", v)
	if err != nil {
		return nil, err
	}

	var er EnrollResponse
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &er, d.Decode(&er)
}

func (ab *AimBrain) actualPostRequest(endpoint string, params map[string]interface{}) ([]byte, error) {
	meth := "POST"
	byt, _ := json.Marshal(params)
	req, err := http.NewRequest(meth, aimbrainAPI+endpoint, bytes.NewBuffer(byt))
	if err != nil {
		return nil, err
	}

	by := bytes.NewBuffer([]byte{})
	err = json.Indent(by, byt, "", "\t")
	if err != nil {
		return nil, err
	}

	sha := sha256.New()
	sha.Write([]byte(ab.ApiSecret))
	h := hmac.New(sha256.New, sha.Sum(nil))
	h.Write([]byte(meth + "\n" + endpoint + "\n" + string(byt)))
	signature := h.Sum(nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-aimbrain-apikey", ab.ApiKey)
	req.Header.Set("X-aimbrain-signature", base64.StdEncoding.EncodeToString(signature))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("aimbrain returned: ", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
