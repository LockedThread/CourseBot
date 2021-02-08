package main

import (
	"github.com/diamondburned/arikawa/v2/discord"
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

	SetupHandlers()

	// Add the needed Gateway intents.
	BotSession.Gateway.AddIntents(gateway.IntentGuildMessages)
	BotSession.Gateway.AddIntents(gateway.IntentDirectMessages)
	BotSession.Gateway.AddIntents(gateway.IntentGuildMessageReactions)

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

func SetupHandlers() {
	BotSession.AddHandler(func(event *gateway.MessageCreateEvent) {
		if len(event.Content) > 0 {
			prefix := event.Content[0]
			if prefix == '!' {
				var index int
				if i := strings.Index(event.Content, " "); i == -1 {
					index = len(event.Content)
				} else {
					index = i + 1
				}
				log.Println("Index: ", index)
				command := event.Content[1:index]
				log.Println("Command: ", command)
				if len(command) > 0 {
					args := strings.Split(event.Content[index:len(event.Content)], " ")
					var input = &CommandInput{
						Command:   command,
						Arguments: args,
						Prefix:    prefix,
						Event:     *event,
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

	BotSession.AddHandler(func(event *gateway.GuildMemberAddEvent) {
		roles, err := BotSession.Roles(event.GuildID)
		HandleErr(err)
		var role discord.Role
		for _, role = range roles {
			if role.Name == "Unverified" {
				break
			}
		}
		err = BotSession.AddRole(event.GuildID, event.User.ID, role.ID)
		HandleErr(err)
	})

	BotSession.AddHandler(func(event *gateway.MessageReactionAddEvent) {
		channel, err := BotSession.Channel(event.ChannelID)
		HandleErr(err)
		if channel.Name == "welcome" {
			log.Println("emoji=",
				event.Emoji)
			// TODO: Switch users role from unverified
		}
	})
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln("We've encountered an error:", err)
	}
}
