// Package api deals with handling the REST requests, communcating between the web
// component and the user messages store
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"

	store "github.com/tomjaroszewskiwork/go-user-messages/app/store"
)

// AddMessage adds a users message to the store.
func AddMessage(w http.ResponseWriter, r *http.Request) {
	var messageBody MessageBody

	// Message validation
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&messageBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if messageBody.Message == "" {
		http.Error(w, "Please send a request body with a message", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	userID := vars["userId"]
	userMessage, err := store.AddMessage(userID, messageBody.Message)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Saving failed", http.StatusInternalServerError)
		return
	}

	// Returns the location of the newly created message
	newLocation := fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, userMessage.MessageID)
	w.Header().Set("Location", path.Clean(newLocation))
	w.WriteHeader(http.StatusCreated)
}

// DeleteMessage deletes a users message from the store.
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	userMessage, err := getMessageEntity(w, r)
	if err != nil || userMessage == nil {
		return
	}
	err = store.DeleteMessage(userMessage)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetFunFacts gets some fun facts about a message!
func GetFunFacts(w http.ResponseWriter, r *http.Request) {
	userMessage, err := getMessageEntity(w, r)
	if err != nil || userMessage == nil {
		return
	}
	funFacts := NewFuncFacts(userMessage.Message)
	bytes, err := json.Marshal(funFacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

// GetMessage gets a specific message for the user.
func GetMessage(w http.ResponseWriter, r *http.Request) {
	userMessage, err := getMessageEntity(w, r)
	if err != nil || userMessage == nil {
		return
	}
	bytes, err := json.Marshal(userMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

// GetMessageList get a list of messages order by generation date, pageinated.
func GetMessageList(w http.ResponseWriter, r *http.Request) {
	// Parses out the parameters
	pageString := r.URL.Query().Get("page")
	// Defaults to first page
	if pageString == "" {
		pageString = "0"
	}
	page, err := strconv.Atoi(pageString)
	if err != nil || page < 0 {
		http.Error(w, "Invalid page parameter, must be a positive integer", http.StatusBadRequest)
		return
	}

	sizeString := r.URL.Query().Get("size")
	// Defaults to 50
	if sizeString == "" {
		sizeString = "50"
	}
	size, err := strconv.Atoi(sizeString)
	if err != nil || size < 1 || size > 100 {
		http.Error(w, "Invalid size parameter, must be a integer between 1 and 100", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	userID := vars["userId"]

	messageEntities, err := store.GetMessages(userID, page, size)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Pulling messages from store failed", http.StatusInternalServerError)
		return
	}
	messageList := NewUserMessageList(messageEntities, size)
	bytes, err := json.Marshal(messageList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func getMessageEntity(w http.ResponseWriter, r *http.Request) (*store.UserMessage, error) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	messageIDtring := vars["messageId"]
	messageID, err := strconv.ParseInt(messageIDtring, 10, 64)
	if err != nil {
		http.Error(w, "Invalid message id", http.StatusBadRequest)
		return nil, err
	}
	userMessage, err := store.GetMessage(userID, messageID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Pulling store failed", http.StatusInternalServerError)
		return nil, err
	}
	if userMessage.MessageID == 0 {
		http.Error(w, "Message not found for that user id", http.StatusNotFound)
		return nil, err
	}
	return userMessage, err
}
