package main

import (
	"fmt"
	"log"
	// "github.com/google/uuid"
)

type Notification struct {
	Id                 int               `json:"id"`
	Title              string            `json:"title"`
	Message            string            `json:"message"`
	DateMsg            string            `json:"date"`
	Topic              TopicNotification `json:"topic"`
	ActiveNotification bool              `json:"active"`
}

type TopicNotification struct {
	TopicId       int    `json:"id"`
	TopicTitle    string `json:"title"`
	TopicInterest bool   `json:"interest"`
}

type Notifications struct {
	Notifications []Notification `json:"notifications"`
}


func ValidNotification(n Notification) {

	if n.ActiveNotification == false {
		log.Fatalln("❌ | Error! To send notification, active permission is required")
	} else {
		fmt.Println("✅ | Permission to send message validated successfully")
	}

}

func ValidTopicNotification(t TopicNotification) {

	if t.TopicInterest == false {
		log.Fatalln("❌ | Error! To send notification, user interest is required")
	} else {
		fmt.Println("✅ | Permission to send message validated successfully")
	}

}

func validAll() {

	notific := Notification{
		Id:      1,
		Title:   "Dragon Ball",
		Message: "Acabou de sair um Post novo",
		DateMsg: "13-09-2023",
		Topic: TopicNotification{
			TopicId:       1,
			TopicTitle:    "Dragon Ball",
			TopicInterest: false,
		},
		ActiveNotification: true,
	}

	ValidNotification(notific)
	ValidTopicNotification(notific.Topic)
}


func main() {
	validAll()
}