package store

import "fmt"

// GetMessage pulls the message from the store
func GetMessage(messageID int64) *UserMessage {
	var userMessage UserMessage
	db.First(&userMessage, messageID)
	// TOOD add user check
	return &userMessage
}

// AddMessage adds a message to the store
func AddMessage(message string) *UserMessage {
	userMessage := UserMessage{Message: message}
	db.Debug().Save(&userMessage)
	fmt.Println(userMessage.MessageID)
	return &userMessage
}