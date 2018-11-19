package Chit_Chat_Server

import (
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func ChitChat(w http.ResponseWriter, r *http.Request) {
}

func LikeMessage(w http.ResponseWriter, r *http.Request) {

}

func DislikeMessage(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" || !strings.Contains(email, urlSubString) {
		json.NewEncoder(w).Encode(BAD_EMAIL)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	hashed, err := bcrypt.GenerateFromPassword(id.NodeID(), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := &User{}

	err = server.DB(DATABASE).C(USERCOLECTION).Find(email).All(user)
	if err != nil {
		user.Email = email
		user.Password = string(hashed)
		err := server.DB(DATABASE).C(USERCOLECTION).Insert(user)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode(FAILED_TO_MODIFY)
		}
	}
	user.Password = string(hashed)
	err = server.DB(DATABASE).C(USERCOLECTION).Update(email, user)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(FAILED_TO_MODIFY)
	}

	json.NewEncoder(w).Encode(SUCCESS)
}
