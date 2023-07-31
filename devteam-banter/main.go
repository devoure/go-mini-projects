package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(">>> Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

var teamMembers = []string{"BAKARI", "ATHUMANI", "DERROW", "JONES", "CLINTON"}
var teamBanters []string

func main() {
	// set random banters
	teamBanters = append(teamBanters, "cant center a div to save his life")
	teamBanters = append(teamBanters, "is slow but the computer is fast")
	teamBanters = append(teamBanters, ", programming is like sex, one mistake and you have to supoort it for the rest of your life")
	teamBanters = append(teamBanters, "uses the word algorithm when they cannot explain what they did")

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5657230133636-5665609479458-VgLbtAAO2CuNvJ17OkJJJnXE")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05KKGEHFM0-5689262948768-46290ce59b2a891b5ab4ec429fbbaa93db60df29c82f5135caf4a35f5368ecdb")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("banter for <team_member>", &slacker.CommandDefinition{
		Description: "team banter slack bot",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			teamMember := strings.ToUpper(request.Param("team_member"))
			botResponse := ""
			if teamMember == "team" {
				randomTeamMember := teamMembers[rand.Intn(len(teamMembers))]
				banter := teamBanters[rand.Intn(len(teamBanters))]
				botResponse = fmt.Sprintf("%s %s", randomTeamMember, banter)
			} else {
				containTeamMember := false
				for _, s := range teamMembers {
					if s == teamMember {
						containTeamMember = true
					}
				}

				if containTeamMember {
					banter := teamBanters[rand.Intn(len(teamBanters))]
					botResponse = fmt.Sprintf("%s %s", teamMember, banter)
				} else {
					botResponse = fmt.Sprintf("%s is not part of team.", teamMember)

				}
			}
			response.Reply(botResponse)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
