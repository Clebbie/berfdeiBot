package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"birthdaybot/util"
)

func main() {
	config, err := util.ReadTokenFromFile("conf.yaml")
	if err != nil {
		fmt.Printf("failed to read config file. %v\n", err.Error())
		return
	}
	discordClient, err := discordgo.New("Bot " + config.AuthToken)

	if err != nil {
		fmt.Printf("failed to start discord bot with auth token %v\n", err.Error())
		return
	}

	discordClient.AddHandler(messageCreate)

	discordClient.Identify.Intents = discordgo.IntentsGuildMessages
	err = discordClient.Open()
	if err != nil {
		fmt.Printf("failed opening connection. %v\n", err.Error())
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	blockUntilSignalReceived()
	_ = discordClient.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	switch message.Content {
	case "ping":
		_, _ = session.ChannelMessageSend(message.ChannelID, "pong")
	case "pong":
		_, _ = session.ChannelMessageSend(message.ChannelID, "ping")
	default:
		return
	}
}

func blockUntilSignalReceived() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChannel
}
