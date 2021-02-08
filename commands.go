package main

import (
	"github.com/diamondburned/arikawa/v2/api"
	"log"
)

/*
Hello and welcome to the {course} discord server. If you're joining one of the course discords for the first time, welcome; if not, welcome back! First and foremost, this discord is in no way directly administered, ran, or affiliated with Valencia College or any other college. We expect everyone to follow common courtesy and respect all others in the discord. By reacting to this message with :white_check_mark: , you are stating you will follow this discord's rules.

If you have any further questions feel free to message Warren#2962.
*/

func CommandSetup(input CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		guild, err := BotSession.Guild(input.Event.GuildID)
		HandleErr(err)
		// TODO: Write code for setup
		role, err := BotSession.CreateRole(guild.ID, api.CreateRoleData{
			Name:        "Unverified",
			Permissions: 0,
			Color:       16187136, // #f6ff00
			Hoist:       true,
			Mentionable: false,
		})
		HandleErr(err)
		roleId := role.ID.String()
		log.Println("roleId=", roleId) // TODO: Put this in database

		/*channels, err := BotSession.Channels(guild.ID)
		HandleErr(err)
		for i := range channels {
			err := BotSession.DeleteChannel(channels[i].ID)
			HandleErr(err)
		}

		_, err = BotSession.CreateChannel(guild.ID, api.CreateChannelData{
			Name:  "general",
			Type:  4,
			Topic: "General communication happens here",
			NSFW:  false,
		})

		HandleErr(err)*/
	}
}

func CommandGetPermissions(input *CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		log.Println(input.Event.ChannelID)
		log.Println(BotSession)
		channel, err := BotSession.Channel(input.Event.ChannelID)
		HandleErr(err)
		log.Println("channel.Permissions=", channel.Permissions)
		for i := range channel.Permissions {
			overwrite := channel.Permissions[i]
			log.Printf("%s %v %v %v", overwrite.ID.String(), overwrite.Allow, overwrite.Deny, overwrite.Type)
		}
	}
}
