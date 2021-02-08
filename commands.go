package main

import (
	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
	"log"
)

func CommandSetup(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		guild, err := BotSession.Guild(input.Event.GuildID)
		HandleErr(err)
		// TODO: Write code for setup
		role, err := BotSession.CreateRole(guild.ID, api.CreateRoleData{
			Name:        "Unverified",
			Permissions: 67174465,
			Color:       16187136, // #f6ff00
			Hoist:       true,
			Mentionable: false,
		})
		HandleErr(err)
		roleId := role.ID.String()
		log.Println("roleId=", roleId) // TODO: Put this in database

		channels, err := BotSession.Channels(guild.ID)
		HandleErr(err)
		for i := range channels {
			err := BotSession.DeleteChannel(channels[i].ID)
			HandleErr(err)
		}

		var permissions = []discord.Overwrite{
			discord.Overwrite{ // The bot
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteMember,
			},
			discord.Overwrite{ // @everyone
				ID:    808402678790225920,
				Allow: 0,
				Deny:  1024,
				Type:  discord.OverwriteRole,
			},
			discord.Overwrite{ // Unverified role
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteRole,
			},
		}

		channel, err := BotSession.CreateChannel(guild.ID, api.CreateChannelData{
			Name:        "welcome",
			Type:        0,
			Topic:       "Confirm acceptance into the discord by reacting to the message.",
			NSFW:        false,
			Permissions: permissions,
		})
		HandleErr(err)

		message, err := BotSession.SendMessage(
			channel.ID,
			"Hello and welcome to the {course} discord server. If you're joining one of the course discords for the first time, welcome; if not, welcome back! First and foremost, this discord is in no way directly administered, ran, or affiliated with Valencia College or any other college. We expect everyone to follow common courtesy and respect all others in the discord. By reacting to this message with :white_check_mark: , you are stating you will follow this discord's rules."+
				"\n\nIf you have any further questions feel free to message this bot.",
			nil,
		)
		HandleErr(err)
		err = BotSession.React(channel.ID, message.ID, ":white_check_mark:")
		HandleErr(err)
	}
}

func CommandGetPermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		channel, err := BotSession.Channel(input.Event.ChannelID)
		HandleErr(err)
		log.Println("channel.Permissions=", channel.Permissions)
		for i := range channel.Permissions {
			overwrite := channel.Permissions[i]
			log.Printf("%s %v %v %v", overwrite.ID.String(), overwrite.Allow, overwrite.Deny, overwrite.Type)
		}
	}
}

func CommandGetRolePermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		roles, err := BotSession.Roles(input.Event.GuildID)
		HandleErr(err)
		for i := range roles {
			if roles[i].ID.String() == input.Arguments[0] {
				permissions := roles[i].Permissions
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
		for i := range channels {
			if channels[i].ID.String() == input.Arguments[0] {
				permissions := channels[i].Permissions
				for i2 := range permissions {
					overwrite := permissions[i2]
					log.Printf("%s %v %v %v", overwrite.ID.String(), overwrite.Allow, overwrite.Deny, overwrite.Type)
				}
				break
			}
		}
	}
}
