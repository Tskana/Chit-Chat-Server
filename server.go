package Chit_Chat_Server

import "gopkg.in/mgo.v2"

type Server struct {
	*mgo.Session
}

const (
	DATABASE          = "chitchat"
	USERCOLECTION     = "chitchatusers"
	MESSAGESCOLECTION = "Messages"
)

var server *Server

func NewServer() {
	server = &Server{}
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}

	server.Session = session
}
