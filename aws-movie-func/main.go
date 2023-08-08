package main

import (
	"math/rand"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

var romanceMovies = []string{"My Sassy Girl", "Love & Other Drugs", "Silver Lining PlayBook"}
var actionMovies = []string{"Guardians of The Galaxy vol 3", "Edge of Tomorrow", "Red"}

func movieRecommeder(genre string) string {
	var movie string
	if genre == "action" {
		movie = romanceMovies[rand.Intn(len(actionMovies))]
	}
	if genre == "romance" {
		movie = actionMovies[rand.Intn(len(romanceMovies))]
	}
	if genre != "action" && genre != "romance" {
		movie = "Error getting movie, check request"
	}

	return movie
}

type FuncEvent struct {
	Genre string `json:"genre"`
}

type FuncResponse struct {
	Movie string `json:"movie"`
}

func HandleLambdaEvent(event FuncEvent) (FuncResponse, error) {
	response := movieRecommeder(strings.ToLower(event.Genre))
	return FuncResponse{Movie: response}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
