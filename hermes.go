package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/matcornic/hermes/v2"
)

func generate( /*username string*/ ) string /*,string*/ {
	//randomly generate a token
	link := "https://jordandde.com/confirm?token=" //+ token
	h := hermes.Hermes{

		Product: hermes.Product{
			Name: "Tugolo",
			Link: "https://tugolo.com",
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jordan",
			Intros: []string{
				"Welcome to Tugolo!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To validate your email, please click here",
					Button: hermes.Button{
						Color: "#ff6347",
						Text:  "Verify your account",
						Link:  link,
					},
				},
			},
		},
	}
	emailBody, err := h.GenerateHTML(email)

	if err != nil {
		panic(err)
	}

	return emailBody /*,link*/
}

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)

	m := mg.NewMessage(
		"Excited User <mailgun@"+domain+">",
		"Hello",
		"",
		"desouza.jordan@yahoo.ca",
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	html := generate()
	m.SetHtml(html)
	_, id, err := mg.Send(ctx, m)
	fmt.Println(err)
	return id, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	domain := os.Getenv("DOMAIN")
	apiKey := os.Getenv("API_key")

	id, er := SendSimpleMessage(domain, apiKey)
	if er != nil {
		fmt.Println(id)
		panic(err)
	}
}
