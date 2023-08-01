package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

var memesArray = []string{"./images/memes/meme1.png",
	"./images/memes/memes2.jpg",
	"./images/memes/memes3.jpg",
	"./images/memes/memes4.jpg",
	"./images/memes/memes5.jpg",
	"./images/memes/memes6.jpg",
}

func main() {
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := []string{os.Getenv("CHANNEL_ID")}

	params := slack.FileUploadParameters{
		Channels: channelArray,
		// random meme from the list
		File: memesArray[rand.Intn(len(memesArray))],
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sent file with details: \n Name: %s \n URL: %s", file.Name, file.URL)
}
