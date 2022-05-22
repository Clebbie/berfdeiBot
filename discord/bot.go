package bday

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"birthdaybot/util"
)

// Bot is the BerfDei struct the container will interact with
type Bot struct {
	Config   *util.Config
	Session  *discordgo.Session
	Commands map[string]func(*discordgo.Session, *discordgo.MessageCreate)
}

// New returns an instance of the Bot.
// There should only ever be one Bot per container
func New() Bot {
	return Bot{
		Config:   nil,
		Session:  nil,
		Commands: map[string]func(*discordgo.Session, *discordgo.MessageCreate){},
	}
}

// Start Begins the discord bot.
// First it loads the configs
// Second creates the sessions
// Adds the message create handlers
// finally it blocks until the os kills the program
func (b *Bot) Start() error {
	if b.Config == nil {
		b.loadConfig()
	}
	if b.Session == nil {
		b.createSession()
	}
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages
	b.addOnMessageCreate()
	err := b.Session.Open()
	if err != nil {
		return fmt.Errorf("failed opening connection. %v\n", err.Error())
	}
	b.blockUntilSignalReceived()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	_ = b.Session.Close()
	return nil
}

// loadConfig reads in the yaml config file and creates the Config struct
func (b *Bot) loadConfig() {
	var err error
	b.Config, err = util.ReadTokenFromFile("conf.yaml")
	if err != nil {
		fmt.Printf("failed to read config file. %v\n", err.Error())
		return
	}
}

// createSession creates the discord session
func (b *Bot) createSession() {
	var err error
	b.Session, err = discordgo.New("Bot " + b.Config.AuthToken)
	if err != nil {
		fmt.Printf("failed to create session %v\n", err.Error())
	}
}

// blockUntilSignalReceived blocks until the program is ended.
func (b *Bot) blockUntilSignalReceived() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChannel
}

// AddCommand adds a command to the list of available commands
func (b *Bot) AddCommand(command string, execution func(*discordgo.Session, *discordgo.MessageCreate)) {
	b.Commands[strings.ToLower(command)] = execution
}

// addOnMessageCreate adds the on message create handler
// First, it checks to see if the message was from the bot
// next, it checks to see if the message is a command and executes if it is
func (b *Bot) addOnMessageCreate() {
	b.Session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}
		if strings.HasPrefix(message.Content, b.Config.CommandPrefix) {
			command := strings.Replace(message.Content, "!", "", -1)
			execution, ok := b.Commands[command]
			if !ok {
				// TODO: Send error response to chat
				return
			}
			execution(session, message)
		}
		switch message.Content {
		case "ping":
			_, _ = session.ChannelMessageSend(message.ChannelID, "pong")
		case "pong":
			_, _ = session.ChannelMessageSend(message.ChannelID, "ping")
		default:
			return
		}
	})
}
