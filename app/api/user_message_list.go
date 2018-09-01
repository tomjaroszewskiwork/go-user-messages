package api

import store "github.com/tomjaroszewskiwork/go-user-messages/app/store"

// UserMessageList is a list of user messages
type UserMessageList struct {

	// If there are more messages for the user to pull
	HasMore bool `json:"hasMore"`

	// List of user messages
	Messages []store.UserMessage `json:"messages"`
}

// NewUserMessageList builds a new list object
func NewUserMessageList(messages []store.UserMessage, requestSize int) *UserMessageList {
	// Checks if there is more messages
	hasMore := false
	if len(messages) > requestSize {
		hasMore = true
		messages = messages[0 : len(messages)-1]
	}
	return &UserMessageList{HasMore: hasMore, Messages: messages}
}
