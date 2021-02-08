package main

import (
	"fmt"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

type CommandInput struct {
	Command   string
	Prefix    uint8
	Arguments []string
	Event     gateway.MessageCreateEvent
}

func (input CommandInput) String() string {
	return fmt.Sprintf(
		"Command=%s, Prefix=%d, Arguments=%s, Event=%v",
		input.Command,
		input.Prefix,
		input.Arguments,
		input.Event)
}

func (input CommandInput) sendMessage(message string) *discord.Message {
	sendMessage, err := BotSession.SendMessage(input.Event.ChannelID, message, nil)
	HandleErr(err)
	return sendMessage
}
