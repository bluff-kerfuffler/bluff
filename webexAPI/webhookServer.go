package webexAPI

import (
	"encoding/json"
	"fmt"
	"strings"
)

type IncomingWebhookData struct {
	Id       string
	Name     string
	Resource string
	Filter   string
	OrgId    string
	AppId    string
	OwnedBy  string
	Status   string
	ActorId  string
	Data     json.RawMessage
}

func handleRawUpdate(data *IncomingWebhookData) {
	switch data.Resource {
	case "messages":
		var message Message
		err := json.Unmarshal(data.Data, &message)
		if err != nil {
			fmt.Println("failed to unmarshal webhook update json", err)
			return
		}
		handleMessageUpdate(data, &message)
	}
}

func handleMessageUpdate(data *IncomingWebhookData, message *Message) {
	words := strings.Fields(message.Text) // could this also be message.markdown?
	if len(words) <= 0 {
		return
	}

	// TODO: proper dispatcher handler instead of bs string comparisons
	switch strings.ToLower(words[0]) {
	case "/secadd":
		secureAddition(message.MentionedPeople...)
	}
}

func secureAddition(people ...string) {
	// TODO: send this to aimbrains
}
