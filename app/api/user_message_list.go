package api

import store "github.com/tomjaroszewskiwork/go-user-messages/app/store"

// UserMessageList is a list of user messages
type UserMessageList struct {

	// If there are more messages for the user to pull
	HasMore bool `json:"hasMore"`

	// List of user messages
	Messages []store.UserMessage `json:"messages"`
}
