package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route holds a specific REST path defintions
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes all paths for the app
type Routes []Route

// NewRouter connects the paths in the app
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// REST call to method mapping
var routes = Routes{
	Route{
		"AddMessage",
		strings.ToUpper("Post"),
		"/v1/users/{userId}/messages",
		AddMessage,
	},

	Route{
		"DeleteMessage",
		strings.ToUpper("Delete"),
		"/v1/users/{userId}/messages/{messageId}",
		DeleteMessage,
	},

	Route{
		"GetFunFacts",
		strings.ToUpper("Get"),
		"/v1/users/{userId}/messages/{messageId}/fun-facts",
		GetFunFacts,
	},

	Route{
		"GetMessage",
		strings.ToUpper("Get"),
		"/v1/users/{userId}/messages/{messageId}",
		GetMessage,
	},

	Route{
		"GetMessageList",
		strings.ToUpper("Get"),
		"/v1/users/{userId}/messages",
		GetMessageList,
	},
}
