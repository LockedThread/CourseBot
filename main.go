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

	s, err := session.New("Bot " + token)
	if err != nil {
		log.Fatalln("Session failed:", err)
	}
	BotSession = *s

	BotSession.AddHandler(func(c *gateway.MessageCreateEvent) {
		if len(c.Content) > 0 {
			prefix := c.Content[0]
			if prefix == '!' {
				var index int
				if i := strings.Index(c.Content, " "); i == -1 {
					index = len(c.Content)
				} else {
					index = i + 1
				}
				log.Println("Index: ", index)
				command := c.Content[1:index]
				log.Println("Command: ", command)
				if len(command) > 0 {
					args := strings.Split(c.Content[index:len(c.Content)], " ")
					var input = &CommandInput{
						Command:   command,
						Arguments: args,
						Prefix:    prefix,
						Event:     *c,
					}
					log.Println("input=", input.String())
					switch strings.TrimSpace(strings.ToLower(command)) {
					case "setup":
						CommandSetup(input)
						break
					case "get_permissions":
						CommandGetPermissions(input)
						break
					case "get_role_permissions":
						CommandGetRolePermissions(input)
						break
					case "get_channel_permissions":
						CommandGetChannelPermissions(input)
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

func HandleErr(err error) {
	if err != nil {
		log.Fatalln("We've encountered an error:", err)
	}
}
