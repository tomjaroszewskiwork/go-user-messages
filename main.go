/*
 * User Messages
 *
 * API for storing and retrieving messages for a user
 *
 * API version: v1
 */

package main

import (
	"log"
	"net/http"

	api "github.com/tomjaroszewskiwork/go-user-messages/app/api"
	store "github.com/tomjaroszewskiwork/go-user-messages/app/store"
)

func main() {
	log.Printf("Server started")

	store.InitDB()

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
