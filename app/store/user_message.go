package store

import (
	"time"
)

// UserMessage message for a user
type UserMessage struct {

	//The time the message was stored at
	GeneratedAt time.Time `json:"generatedAt"`

	// Message content
	Message string `json:"message" gorm:"not_null"`

	// Message id
	MessageID int64 `json:"messageId" gorm:"PRIMARY_KEY"`

	// User id
	UserID string `json:"userId" gorm:"index:userIdx"`
}
