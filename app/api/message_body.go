package api

// Passed in when creating a new message
type MessageBody struct {

	// The new message to be stored
	Message string `json:"message"`
}
