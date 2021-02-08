package main

import "log"

func CommandSetup(input CommandInput) {
	if input.Event.Author.ID.String() == "697631712485572648" {
		guild, err := BotSession.Guild(input.Event.GuildID)
		if err != nil {
			log.Fatalf("We've got a problem here: %d", err)
		}
		// TODO: Write code for setup
	}
}
