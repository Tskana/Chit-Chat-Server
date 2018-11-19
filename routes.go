package Chit_Chat_Server

import (
	"net/http"
)

type Route struct {
	name    string
	method  string
	pattern string
	http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"chitchat get",
		"GET",
		"/chitchat",
		ChitChat,
	},
	Route{
		"chitchat post",
		"POST",
		"/chitchat",
		ChitChat,
	},
	Route{
		"like",
		"POST",
		"/chitchat/like/<post_id>",
		LikeMessage,
	},
	Route{
		"dislike",
		"GET",
		"/chitchat/dislike/<post_id>",
		DislikeMessage,
	},
	Route{
		"register",
		"GET",
		"/chitchat/register",
		Register,
	},
}
