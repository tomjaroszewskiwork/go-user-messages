package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter()
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
