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
	BotSession    session.Session
	GuildCacheMap = make(map[discord.GuildID]GuildCache)
)

func main() {
	if token := os.Getenv("BOT_TOKEN"); len(token) == 0 {
		log.Fatalln("No BOT_TOKEN environmental variable given. Correct this in the docker env settings.")
	} else {
		s, err := session.New("Bot " + token)
		if err != nil {
			log.Fatalln("Session failed:", err)
		}
		BotSession = *s
	}

	SetupHandlers()

	// Add the needed Gateway intents.
	BotSession.Gateway.AddIntents(gateway.IntentGuildMessages)
	BotSession.Gateway.AddIntents(gateway.IntentDirectMessages)
	BotSession.Gateway.AddIntents(gateway.IntentGuildMessageReactions)

	if err := BotSession.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer BotSession.Close()

	log.Println("CourseBot started without issues")
	LoadGuildCache()
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
				command := event.Content[1:index]
				if len(command) > 0 {
					args := strings.Split(event.Content[index:len(event.Content)], " ")
					var input = &CommandInput{
						Command:   command,
						Arguments: args,
						Prefix:    prefix,
						Event:     *event,
					}
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
		err := BotSession.AddRole(event.GuildID, event.User.ID, GuildCacheMap[event.GuildID].UnverifiedID)
		HandleErr(err)
	})

	BotSession.AddHandler(func(event *gateway.MessageReactionAddEvent) {
		if event.Member.User.Bot {
			return
		}
		channel, err := BotSession.Channel(event.ChannelID)
		HandleErr(err)
		if channel.Name == "welcome" {
			cache := GuildCacheMap[event.GuildID]
			err := BotSession.RemoveRole(event.GuildID, event.UserID, cache.UnverifiedID)
			HandleErr(err)
			err = BotSession.AddRole(event.GuildID, event.UserID, cache.MemberID)
			HandleErr(err)
		}
	})
}

func LoadGuildCache() {
	guilds, err := BotSession.Guilds(0)
	HandleErr(err)
	for _, guild := range guilds {
		AddGuildCache(guild.ID)
	}
}

func AddGuildCache(guildId discord.GuildID) {
	channels, err := BotSession.Channels(guildId)
	HandleErr(err)
	var welcomeChannelId discord.ChannelID
	for _, channel := range channels {
		if channel.Name == "welcome" {
			welcomeChannelId = channel.ID
			break
		}
	}
	roles, err := BotSession.Roles(guildId)
	HandleErr(err)
	var unverifiedId, memberId discord.RoleID
	for _, role := range roles {
		if role.Name == "Unverified" {
			unverifiedId = role.ID
		} else if role.Name == "Member" {
			memberId = role.ID
		}
	}
	guild, err := BotSession.Guild(guildId)
	HandleErr(err)
	GuildCacheMap[guildId] = GuildCache{
		GuildID:          guildId,
		WelcomeChannelID: welcomeChannelId,
		UnverifiedID:     unverifiedId,
		MemberID:         memberId,
		CourseName:       guild.Name,
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Panicln("We've encountered an error:", err)
	}
}
