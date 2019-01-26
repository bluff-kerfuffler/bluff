package webexAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type User struct {
	Id           string
	Emails       []string
	DisplayName  string
	NickName     string
	FirstName    string
	LastName     string
	Avatar       string
	OrgId        string
	Roles        []string
	Licenses     []string
	Created      string
	Timezone     string
	LastActivity string
	// Status
	// InvitePending
	// LoginEnabled
}

func (b Bot) GetMe() (*User, error) {
	res, err := b.Get("people/me", url.Values{})
	if err != nil {
		return nil, err
	}

	var u User
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &u, d.Decode(&u)
}

func (b Bot) GetUser(userId string) (*User, error) {
	v := url.Values{}
	v.Set("personId", userId)
	res, err := b.Get("people/"+userId, v)
	if err != nil {
		return nil, err
	}

	var u User
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &u, d.Decode(&u)
}

type RoomResponse struct {
	Items []Room
}

type Room struct {
	Id    string
	Title string
	//Type enum
	IsLocked     bool
	TeamId       string
	LastActivity string
	CreatorId    string
	Created      string
}

func (b Bot) GetRooms() (*RoomResponse, error) {
	v := url.Values{}
	res, err := b.Get("rooms", v)
	if err != nil {
		return nil, err
	}

	var rr RoomResponse
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &rr, d.Decode(&rr)
}

type MessageResponse struct {
	Items []Message
}

type Message struct {
	Id     string
	RoomId string
	//RoomType enum // group/direct
	ToPersonId      string // only in direct
	ToPersonEmail   string // only in direct
	Text            string
	Markdown        string
	Files           []string
	PersonId        string
	MentionedPeople []string
	MentionedGroups []string
	Created         string
}

func (b Bot) GetMessages(roomId string, nMessages int) (*MessageResponse, error) {
	v := url.Values{}
	v.Set("roomId", roomId)
	v.Set("max", strconv.Itoa(nMessages))
	res, err := b.Get("messages", v)
	if err != nil {
		return nil, err
	}

	var mr MessageResponse
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &mr, d.Decode(&mr)
}

func (b Bot) SendMessage(roomId string, text string) (*Message, error) {
	v := map[string]interface{}{
		"roomId": roomId,
		"text":   text,
	}
	res, err := b.Post("messages", v)
	if err != nil {
		return nil, err
	}

	var m Message
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &m, d.Decode(&m)
}
func (b Bot) SendMessageMarkdown(roomId string, markdown string) (*Message, error) {
	v := map[string]interface{}{
		"roomId":   roomId,
		"markdown": markdown,
	}
	res, err := b.Post("messages", v)
	if err != nil {
		return nil, err
	}

	var m Message
	d := json.NewDecoder(bytes.NewBuffer(res))
	return &m, d.Decode(&m)
}

func MentionIdMarkdown(userId string, name string) string {
	return fmt.Sprintf("<@personId:%s|%s>", userId, name)
}
func MentionEmailMarkdown(email string, name string) string {
	return fmt.Sprintf("<@personEmail:%s|%s>", email, name)
}
