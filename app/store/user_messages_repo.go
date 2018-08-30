package store

import (
	"time"

	gorm "github.com/jinzhu/gorm"
)

// GetMessage pulls the message from the store
func GetMessage(userID string, messageID int64) (*UserMessage, error) {
	var userMessage UserMessage
	err := db.Where("user_id = ? AND message_id = ?", userID, messageID).First(&userMessage).Error
	// Eats the not found to make things cleaner
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	return &userMessage, err
}

// AddMessage adds a message to the store
func AddMessage(userID string, message string) (*UserMessage, error) {
	userMessage := UserMessage{UserID: userID, Message: message, GeneratedAt: time.Now().UTC()}
	err := db.Create(&userMessage).Error
	return &userMessage, err
}
