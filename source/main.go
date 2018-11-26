package main

import (
	chitchat "github.com/tskana/Chit-Chat-Server"
	"log"
	"net/http"
)

const addr = "http://localhost"

func init() {
	chitchat.NewServer()
}

func main() {
	log.Fatal(http.ListenAndServe(addr, chitchat.NewRoutes()))
}
