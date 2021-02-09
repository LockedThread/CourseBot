package main

import (
	"fmt"
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

		memberRoleId := CreateRole(
			guild.ID,
			"Member",
			103926337,
			16764928,
			true,
			false,
		)

		unverifiedRoleId := CreateRole(
			guild.ID,
			"Unverified",
			67174465,
			16187136,
			true,
			false,
		) // Unverified Role

		infoChannelId := CreateChannel(guild.ID, "info", 0, []discord.Overwrite{
			{ // The bot
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteMember,
			},
			{ // Member role
				ID:    discord.Snowflake(memberRoleId),
				Allow: 1024,
				Deny:  2048,
				Type:  discord.OverwriteRole,
			}}, nil,
		)

		mention := "<@697631712485572648>"
		_, err = BotSession.SendMessage(infoChannelId,
			fmt.Sprintf("**Purpose**\nThe purpose of this discord is an attempt to foster the same college culture we would have if we were on campus. Along with that, this discord gives students a better opportunity to create study groups and ask each other questions about assignments.\n\n**What we ask from you**\nWe ask that you do not upload any complete course material directly to this discord so we are not liable for your actions. We also ask that you try to spread the word about this discord and our other discords.\n\n**Want to open another discord?**\nWe're always looking to open new discords for other courses, if you want new one open then just message %s and he can create another one for the course.\n\n**Final Message**\nMy suggestion is that people in this server should plan out times of the week where people can study together; remember, your strength grows in numbers. If you have any questions feel free to message %s as he is the one that runs this discord.\n\n**Disclaimer**\nThis discord is in no way directly affiliated with any college and is solely ran by students for students.", mention, mention), nil)
		HandleErr(err)

		welcomeChannelId := CreateChannel(guild.ID, "welcome", 0, []discord.Overwrite{
			{ // The bot
				ID:    808400745597632562,
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteMember,
			},
			{ // Unverified role
				ID:    discord.Snowflake(unverifiedRoleId),
				Allow: 1024,
				Deny:  0,
				Type:  discord.OverwriteRole,
			}}, nil,
		) // Welcome Text Channel

		message, err := BotSession.SendMessage(
			welcomeChannelId,
			fmt.Sprintf("Hello and welcome to the %s discord server. If you're joining one of the course discords for the first time, welcome; if not, welcome back! First and foremost, this discord is in no way directly administered, ran, or affiliated with Valencia College or any other college. We expect everyone to follow common courtesy and respect all others in the discord. By reacting to this message with :white_check_mark: , you are stating you will follow this discord's rules."+
				"\n\nIf you have any further questions feel free to message <@697631712485572648>.", guild.Name),
			nil,
		) // Welcome Message

		generalPermissions := []discord.Overwrite{{ // Member role
			ID:    discord.Snowflake(memberRoleId),
			Allow: 1024,
			Deny:  2048,
			Type:  discord.OverwriteRole,
		}}

		categoryId := CreateCategory(guild.ID, "general", generalPermissions)

		CreateChannel(guild.ID, "general", 0, generalPermissions, &categoryId)
		CreateChannel(guild.ID, "homework-help", 0, generalPermissions, &categoryId)
		CreateChannel(guild.ID, "off-topic", 0, generalPermissions, &categoryId)

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
