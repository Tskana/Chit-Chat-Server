package Chit_Chat_Server

import (
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func ChitChat(w http.ResponseWriter, r *http.Request) {
	if !Authorize(r.Form) {
		json.NewEncoder(w).Encode(UNAUTHERIZED_USER)
		return
	}

	switch r.Method {
	case "GET":
		skip, err := strconv.Atoi(r.FormValue("skip"))

		if err != nil {
			skip = 0
		}

		limit, err := strconv.Atoi(r.FormValue("limit"))
		if err != nil {
			limit = 20
		}

		messages := make([]Message, limit)

		server.DB(DATABASE).C(MESSAGESCOLECTION).Find(nil).Skip(skip).Batch(limit).All(messages)

		rm := &ChitChatMessages{
			Date:     time.Now().Format("Mon Jan 2 15:04:05 2006"),
			Count:    len(messages),
			Messages: messages,
		}

		json.NewEncoder(w).Encode(rm)
	case "POST":
		client := r.FormValue("client")

		message := r.PostFormValue("message")
		lat := r.PostFormValue("lat")
		lon := r.PostFormValue("lon")

		if message != "" && client != "" {
			id, err := uuid.NewUUID()
			if err != nil {
				log.Fatal(err)
				return
			}
			ms := &Message{
				ID: id.String(),
				Client: client,
				Date: time.Now().Format("Mon Jan 2 15:04:05 2006"),
				IP: r.RemoteAddr,
				Location: []string{lat, lon},
				Message: message,
			}

			server.DB(DATABASE).C(MESSAGESCOLECTION).Insert(ms)
			json.NewEncoder(w).Encode(SUCCESS)
		}
		json.NewEncoder(w).Encode(MISSING_PARAMS)
	default:
		json.NewEncoder(w).Encode(BAD_METHOD)
	}
}

// TODO: Can make like and dislike smaller with less duplicate code
func LikeMessage(w http.ResponseWriter, r *http.Request) {
	if !Authorize(r.Form) {
		json.NewEncoder(w).Encode(UNAUTHERIZED_USER)
	}

	u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := strconv.Atoi(path.Base(u.Path))
	if err != nil {
		log.Fatal(err)
		return
	}

	ms := &Message{}

	err = server.DB(DATABASE).C(MESSAGESCOLECTION).FindId(id).One(ms)
	if err != nil {
		log.Fatal(err)
		return
	}

	ms.Likes += 1

	err = server.DB(DATABASE).C(MESSAGESCOLECTION).UpdateId(id, ms)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(BAD_UPDATE)
		return
	}

	json.NewEncoder(w).Encode(SUCCESS)
}

func DislikeMessage(w http.ResponseWriter, r *http.Request) {
	if !Authorize(r.Form) {
		json.NewEncoder(w).Encode(UNAUTHERIZED_USER)
	}

	u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := strconv.Atoi(path.Base(u.Path))
	if err != nil {
		log.Fatal(err)
		return
	}

	ms := &Message{}

	err = server.DB(DATABASE).C(MESSAGESCOLECTION).FindId(id).One(ms)
	if err != nil {
		log.Fatal(err)
		return
	}

	ms.Dislikes += 1

	err = server.DB(DATABASE).C(MESSAGESCOLECTION).UpdateId(id, ms)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(BAD_UPDATE)
		return
	}

	json.NewEncoder(w).Encode(SUCCESS)
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
