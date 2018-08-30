package store

// GetMessage pulls the message from the store
func GetMessage(messageID int64) *UserMessage {
	var userMessage UserMessage
	db.First(&userMessage, messageID)
	// TOOD add user check
	return &userMessage
}

// AddMessage adds a message to the store
func AddMessage(userID string, message string) *UserMessage {
	userMessage := UserMessage{UserID: userID, Message: message}
	db.Debug().Save(&userMessage)
	return &userMessage
}
