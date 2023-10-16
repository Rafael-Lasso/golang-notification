package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

func mainSMS() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	message := "Acabou de sair um Post novo com um tema que voce vai adorar! Vem ver 😁: https://www.youtube.com/@SpeedUzumaki"

	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom("+15855656329")
	params.SetTo("+5518997635588")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}

}
