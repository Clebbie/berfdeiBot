package main

import (
	"github.com/bwmarrin/discordgo"

	bday "birthdaybot/discord"
)

func main() {
	bot := bday.New()
	bot.AddCommand("greet", func(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Mushi Mushi")
	})
	bot.Start()
}
