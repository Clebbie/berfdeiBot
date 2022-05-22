package discord

import "github.com/bwmarrin/discordgo"

func createClient(authToken string) {
	discord, err := discordgo.New("Bot " + authToken)
}
