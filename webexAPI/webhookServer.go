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

func (b Bot) handleRawUpdate(data *IncomingWebhookData) {
	switch data.Resource {
	case "messages":
		var message Message
		err := json.Unmarshal(data.Data, &message)
		if err != nil {
			fmt.Println("failed to unmarshal webhook update json", err)
			return
		}
		b.handleMessageUpdate(data, &message)
	}
}

func (b Bot) handleMessageUpdate(data *IncomingWebhookData, message *Message) {
	words := strings.Fields(message.Text) // could this also be message.markdown?
	if len(words) <= 0 {
		return
	}

	// TODO: proper dispatcher handler instead of bs string comparisons
	switch strings.ToLower(words[0]) {
	case "/secadd":
		b.secureAddition(data.ActorId, message.MentionedPeople...)
	}
}

func (b Bot) secureAddition(requester string, people ...string) {
	for _, p := range people {
		b.SendPrivateMessageMarkdown(p, MentionIdMarkdown(requester, "This person") + " has invited you to join a secure team.\n\nClick here to join: (how do I get this link)")
	}
}
