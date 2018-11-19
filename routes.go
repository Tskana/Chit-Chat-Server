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
		"get",
		"/chitchat",
		ChitChat,
	},
	Route{
		"chitchat post",
		"post",
		"/chitchat",
		ChitChat,
	},
	Route{
		"like",
		"post",
		"/chitchat/like/<post_id>",
		LikeMessage,
	},
	Route{
		"dislike",
		"post",
		"/chitchat/dislike/<post_id",
		DislikeMessage,
	},
	Route{
		"register",
		"post",
		"/chitchat/register",
		Register,
	},
}
