package Chit_Chat_Server

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/smtp"
)

func SendEmail(content string, email string) error {
	c, err := smtp.Dial(mailUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer c.Close()
	c.Mail(mailUrl)
	c.Rcpt(email)
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer wc.Close()

	buff := bytes.NewBufferString("Subject: Your Chit Chat API Key\r\n" + "\r\n" + content + "\r\n")

	_, err = buff.WriteTo(wc)
	return err
}

func Autherize(client string, key []byte) bool {
	user := &User{}
	err := server.DB(DATABASE).C(USERCOLECTION).Find(client).All(user)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), key)
	if err != nil {
		return false
	}

	return true
}
