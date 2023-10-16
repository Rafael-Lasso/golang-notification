package main
// Download the helper library from https://www.twilio.com/docs/go/install

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func WhatsappService() {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	whatsappFROM := os.Getenv("WHATSAPP_FROM")
	whatsappTO := os.Getenv("WHATSAPP_TO")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}
	params.SetFrom(whatsappFROM)
	params.SetBody("Acabou de sair um Post novo com um tema que voce vai adorar! Vem ver 😁:")
	params.SetTo(whatsappTO)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
			fmt.Println("Message sent successfully")

		} else {
			fmt.Println(resp.Sid)
			fmt.Println("Failed to send message")
		}
	}
}

func main() {
	WhatsappService()
}