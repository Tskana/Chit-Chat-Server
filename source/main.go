package main

import (
	chitchat "github.com/tsukinai/Chit-Chat-Server"
	"log"
	"net/http"
)

const addr = "http://localhost"

func init() {
	chitchat.NewServer()
}

func main() {
	router := chitchat.NewRoutes()

	log.Fatal(http.ListenAndServe(addr, router))
}