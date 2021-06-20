package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

var key = os.Getenv("SMTP_KEY")
var pass = os.Getenv("SMTP_PASS")

func SendResponse(w http.ResponseWriter, data interface{}, code int) {
	b, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(b)
	return
}

func generateID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

var from = &mailjet.RecipientV31{
	Email: "vashishtiv@gmail.com",
	Name:  "Yashi Gupta",
}

var client = mailjet.NewMailjetClient(key, pass)

func SendEmail(RecipientEmail string, RecipientName string, Body string) bool {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: from,
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: RecipientEmail,
					Name:  RecipientName,
				},
			},
			Subject:  "Interview Scheduled",
			TextPart: "You have an Interview Scheduled " + Body,
			HTMLPart: ``,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := client.SendMailV31(&messages)
	if err != nil {
		fmt.Println("error while sending mail", err)
		return false
	}

	fmt.Println(res.ResultsV31)
	return true
}

// func RequestParamCheck(){

// }
