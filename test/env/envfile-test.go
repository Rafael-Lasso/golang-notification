package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func validEnvFile() {

	SEND_TO_NAME := os.Getenv("SEND_TO_NAME")
	SEND_FROM_NAME := os.Getenv("SEND_FROM_NAME")
	SEND_TO_ADDRESS := os.Getenv("SEND_TO_ADDRESS")
	SEND_FROM_ADDRESS := os.Getenv("SEND_FROM_ADDRESS")
	ACCOUNT_SID := os.Getenv("ACCOUNT_SID")
	AUTH_TOKEN := os.Getenv("AUTH_TOKEN")
	TWILIO_TO := os.Getenv("TWILIO_TO")
	TWILIO_FROM := os.Getenv("TWILIO_FROM")

	if SEND_TO_NAME == os.Getenv("SEND_TO_NAME") {
		fmt.Println("✅ | Loaded var successfully: SEND_TO_NAME", SEND_TO_NAME)
	} else {
		log.Fatalln("❌ | Error loading .env var SEND_TO_NAME")
	}

	if SEND_FROM_NAME == os.Getenv("SEND_FROM_NAME") {
		fmt.Println("✅ | Loaded var successfully: SEND_FROM_NAME", SEND_FROM_NAME)
	} else {
		log.Fatalln("❌ | Error loading .env var SEND_FROM_NAME ")
	}

	if SEND_TO_ADDRESS == os.Getenv("SEND_TO_ADDRESS") {
		fmt.Println("✅ | Loaded var successfully: SEND_TO_ADDRESS", SEND_TO_ADDRESS)
	} else {
		log.Fatalln("❌ | Error loading .env var SEND_TO_ADDRESS")
	}

	if SEND_FROM_ADDRESS == os.Getenv("SEND_FROM_ADDRESS") {
		fmt.Println("✅ | Loaded var successfully: SEND_FROM_ADDRESS", SEND_FROM_ADDRESS)
	} else {
		log.Fatalln("❌ | Error loading .env var SEND_FROM_ADDRESS")
	}

	if ACCOUNT_SID == os.Getenv("ACCOUNT_SID") {
		fmt.Println("✅ | Loaded var successfully: ACCOUNT_SID", ACCOUNT_SID)
	} else {
		log.Fatalln("❌ | Error loading .env var ACCOUNT_SID")
	}

	if AUTH_TOKEN == os.Getenv("AUTH_TOKEN") {
		fmt.Println("✅ | Loaded var successfully: AUTH_TOKEN", AUTH_TOKEN)
	} else {
		log.Fatalln("❌ | Error loading .env var AUTH_TOKEN")
	}

	if TWILIO_TO == os.Getenv("TWILIO_TO") {
		fmt.Println("✅ | Loaded var successfully: TWILIO_TO", TWILIO_TO)
	} else {
		log.Fatalln("❌ | Error loading .env var TWILIO_TO")
	}

	if TWILIO_FROM == os.Getenv("TWILIO_FROM") {
		fmt.Println("✅ | Loaded var successfully: TWILIO_FROM", TWILIO_FROM)
	} else {
		log.Fatalln("❌ | Error loading .env var TWILIO_FROM")
	}

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validEnvFile()
}
