package api

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"

	store "github.com/tjaroszewskiwork/go-user-messages/app/store" 
)

// AddMessage adds a users message to the store
func AddMessage(w http.ResponseWriter, r *http.Request) {
	var messageBody MessageBody
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&messageBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userMessage := store.AddMessage(messageBody.Message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path,userMessage.MessageID) )
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
	//userId := vars["userId"]
	messageIDtring := vars["messageId"]
	messageID, err := strconv.ParseInt(messageIDtring, 10, 64)
	if err != nil {
		http.Error(w, "Invalid message id", http.StatusBadRequest)
		return
	}
	userMessage := store.GetMessage(messageID)
	bytes, _ := json.Marshal(userMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetMessageList get a list of messages order by generation date, pageinated
func GetMessageList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
