package main

import (
	"github.com/diamondburned/arikawa/v2/api"
	"github.com/diamondburned/arikawa/v2/discord"
)

func CreateRole(guildId discord.GuildID, roleName string, permissions discord.Permissions, color discord.Color, hoist bool, mentionable bool) discord.RoleID {
	role, err := BotSession.CreateRole(guildId, api.CreateRoleData{
		Name:        roleName,
		Permissions: permissions,
		Color:       color, // #f6ff00
		Hoist:       hoist,
		Mentionable: mentionable,
	})
	HandleErr(err)
	return role.ID
	// TODO: Put this in database
}

func CreateChannel(guildId discord.GuildID, channelName string, channelType discord.ChannelType, permissions []discord.Overwrite) discord.ChannelID {
	channel, err := BotSession.CreateChannel(guildId, api.CreateChannelData{
		Name:        channelName,
		Type:        channelType,
		Permissions: permissions,
	})
	HandleErr(err)
	return channel.ID
}
