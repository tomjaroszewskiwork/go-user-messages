// Package store deals with storing and retriving users messages in a persistant store
package store

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GetMessage pulls the message from the store
func GetMessage(userID string, messageID int64) (*UserMessage, error) {
	var userMessage UserMessage
	err := DB.Where("user_id = ? AND message_id = ?", userID, messageID).First(&userMessage).Error
	// Eats the not found to make things cleaner
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return &userMessage, err
}

// GetMessages gets a list user messages based on the page, size and offset, sorted by generation times
func GetMessages(userID string, page int, size int, buffer int) ([]UserMessage, error) {
	offsetStart := page * size
	offsetEnd := (page+1)*size + buffer
	var userMessages []UserMessage
	err := DB.Where("user_id = ? ", userID).
		Order("generated_at desc").
		Offset(offsetStart).
		Limit(offsetEnd).
		Find(&userMessages).Error
	return userMessages, err
}

// AddMessage adds a message to the store
func AddMessage(userID string, message string) (*UserMessage, error) {
	// Everything is stored as UTC to make things consistant no matter the timezone of the deployment
	userMessage := UserMessage{UserID: userID, Message: message, GeneratedAt: time.Now().UTC()}
	err := DB.Create(&userMessage).Error
	return &userMessage, err
}

// DeleteMessage deletes the given message from the store
func DeleteMessage(userMessage *UserMessage) error {
	err := DB.Delete(&userMessage).Error
	return err
}
