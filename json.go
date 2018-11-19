package Chit_Chat_Server

type Result struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Reason  string `json:"reason"`
}

type User struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Message struct {
	ID       string `json:"_id"`
	Client   string `json:"client"`
	Date     string `json:"date"`
	Dislikes int    `bson:"dislikes"`
	IP       string `json:"ip"`
	Likes    int    `bson:"likes"`
	Location []*int `json:"loc"` // This will likely not work
	Message  string `json:"message"`
}

type ChitChatMessages struct {
	Count    int       `bson:"count"`
	Date     string    `json:"date"`
	Messages []Message `json:"messages"`
}
