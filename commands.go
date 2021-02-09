package main

import (
	"github.com/diamondburned/arikawa/v2/discord"
	"log"
)

func CommandSetup(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		guild, err := BotSession.Guild(input.Event.GuildID)
		HandleErr(err)

		// Wipe Channels Start
		channels, err := BotSession.Channels(guild.ID)
		HandleErr(err)
		for _, channel := range channels {
			err := BotSession.DeleteChannel(channel.ID)
			HandleErr(err)
		}
		// Wipe Channels End

		// Wipe Roles Start
		roles, err := BotSession.Roles(guild.ID)
		HandleErr(err)
		for _, role := range roles {
			if role.Name != "CourseBot" && role.Name != "@everyone" && !role.Managed {
				log.Println("role=", role.Name)
				err := BotSession.DeleteRole(guild.ID, role.ID)
				HandleErr(err)
			}
		}
		// Wipe Roles End

		CreateRole(
			guild.ID,
			"Unverified",
			67174465,
			16187136,
			true, false,
		) // Unverified Role

		CreateRole(
			guild.ID,
			"Member",
			103926337,
			16764928,
			true,
			false,
		) // Member Role

		welcomeChannelId := CreateChannel(guild.ID, "welcome", 0, []discord.Overwrite{
			{ // The bot
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteMember,
			},
			{ // @everyone
				ID:    808402678790225920,
				Allow: 0,
				Deny:  1024,
				Type:  discord.OverwriteRole,
			},
			{ // Unverified role
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteRole,
			}},
		) // Welcome Text Channel

		message, err := BotSession.SendMessage(
			welcomeChannelId,
			"Hello and welcome to the {course} discord server. If you're joining one of the course discords for the first time, welcome; if not, welcome back! First and foremost, this discord is in no way directly administered, ran, or affiliated with Valencia College or any other college. We expect everyone to follow common courtesy and respect all others in the discord. By reacting to this message with :white_check_mark: , you are stating you will follow this discord's rules."+
				"\n\nIf you have any further questions feel free to message this bot.",
			nil,
		) // Welcome Message
		HandleErr(err)
		err = BotSession.React(welcomeChannelId, message.ID, "✅") // Welcome Message Reaction
		HandleErr(err)

		AddGuildCache(guild.ID)
	}
}

func CommandGetPermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		channel, err := BotSession.Channel(input.Event.ChannelID)
		HandleErr(err)
		for _, permission := range channel.Permissions {
			log.Printf(
				"Permission | %s | %v | %v | %v",
				permission.ID.String(),
				permission.Allow,
				permission.Deny,
				permission.Type,
			)
		}
	}
}

func CommandGetRolePermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		roles, err := BotSession.Roles(input.Event.GuildID)
		HandleErr(err)
		for _, role := range roles {
			if role.ID.String() == input.Arguments[0] {
				permissions := role.Permissions
				log.Println("permissions=", permissions)
				break
			}
		}
	}
}

func CommandGetChannelPermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		channels, err := BotSession.Channels(input.Event.GuildID)
		HandleErr(err)
		for _, channel := range channels {
			if channel.ID.String() == input.Arguments[0] {
				for _, permission := range channel.Permissions {
					log.Printf(
						"Permission %s | %v | %v | %v",
						permission.ID.String(),
						permission.Allow,
						permission.Deny,
						permission.Type,
					)
				}
				break
			}
		}
	}
}
