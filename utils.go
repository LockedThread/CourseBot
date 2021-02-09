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

func CreateCategory(guildId discord.GuildID, categoryName string, permissions []discord.Overwrite) discord.ChannelID {
	channel, err := BotSession.CreateChannel(guildId, api.CreateChannelData{
		Name:        categoryName,
		Type:        4,
		Permissions: permissions,
	})
	HandleErr(err)
	return channel.ID
}

func CreateChannel(guildId discord.GuildID, channelName string, channelType discord.ChannelType, permissions []discord.Overwrite, categoryId *discord.ChannelID) discord.ChannelID {
	var id discord.ChannelID
	if categoryId == nil {
		id = 0
	} else {
		id = *categoryId
	}

	channel, err := BotSession.CreateChannel(guildId, api.CreateChannelData{
		Name:        channelName,
		Type:        channelType,
		Permissions: permissions,
		CategoryID:  id,
	})
	HandleErr(err)
	return channel.ID
}
