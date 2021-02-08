package main

import (
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/session"
	"log"
	"os"
	"strings"
)

var (
	BotSession session.Session
)

func main() {
	var token = os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	BotSession, err := session.New("Bot " + token)
	if err != nil {
		log.Fatalln("Session failed:", err)
	}

	BotSession.AddHandler(func(c *gateway.MessageCreateEvent) {
		if len(c.Content) > 0 {
			prefix := c.Content[0]
			if prefix == '!' {
				index := strings.Index(c.Content, " ")
				command := c.Content[0:index]
				if len(command) > 0 {
					args := strings.Split(c.Content[index:len(c.Content)], " ")
					var input = CommandInput{
						Command:   command,
						Arguments: args,
						Prefix:    prefix,
						Event:     c,
					}
					switch strings.ToLower(command) {
					case "setup":
						CommandSetup(input)
						break
					default:
						break
					}
				}
			}
		}
	})

	// Add the needed Gateway intents.
	BotSession.Gateway.AddIntents(gateway.IntentGuildMessages)
	BotSession.Gateway.AddIntents(gateway.IntentDirectMessages)

	if err := BotSession.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer BotSession.Close()

	u, err := BotSession.Me()
	if err != nil {
		log.Fatalln("Failed to get myself:", err)
	}

	log.Println("Started as", u.Username)

	// Block forever.
	select {}
}
