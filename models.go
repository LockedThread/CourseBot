package main

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"log"
)

type CommandInput struct {
	Command   string
	Prefix    uint8
	Arguments []string
	Event     *gateway.MessageCreateEvent
}

func (input *CommandInput) sendMessage(message string) *discord.Message {
	sendMessage, err := BotSession.SendMessage(input.Event.ChannelID, message, nil)
	if err != nil {
		log.Printf("We've got a problem: %d", err)
	}
	return sendMessage
}
