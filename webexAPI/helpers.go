package webexAPI

import "fmt"

func MentionIdMarkdown(userId string, name string) string {
	return fmt.Sprintf("<@personId:%s|%s>", userId, name)
}
func MentionEmailMarkdown(email string, name string) string {
	return fmt.Sprintf("<@personEmail:%s|%s>", email, name)
}
