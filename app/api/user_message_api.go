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

// AddMessage adds a users message to the store
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

// DeleteMessage removes a users message from the store
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetFunFacts gets some fun facts about a message!
func GetFunFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetMessage gets a specific message
func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	messageIDtring := vars["messageId"]
	messageID, err := strconv.ParseInt(messageIDtring, 10, 64)
	if err != nil {
		http.Error(w, "Invalid message id", http.StatusBadRequest)
		return
	}
	userMessage, err := store.GetMessage(userID, messageID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Pulling store failed", http.StatusInternalServerError)
		return
	}
	if userMessage.MessageID == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	bytes, err := json.Marshal(userMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetMessageList get a list of messages order by generation date, pageinated
func GetMessageList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
