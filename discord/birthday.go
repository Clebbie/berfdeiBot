package bday

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Event is a generic type that represents an occasion on a calendar
type Event interface {
	Announcement(*discordgo.Session, *discordgo.MessageCreate)
}

// Birthday is an event specific for peoples birthdays
type Birthday struct {
	PersonName string
	PersonId   string
	Date       time.Time
}

func (b *Birthday) Announcement(session *discordgo.Session, message *discordgo.MessageCreate) {

}

// Calendar is a map of days to an array of events
type Calendar struct {
	Events map[time.Time][]Event
}
