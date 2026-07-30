package commands

import (
	"kessa/handlers"
	"time"

	"github.com/bwmarrin/discordgo"
)

var example = &handlers.Command{
	Name:        "ex",
	Aliases:     []string{"x"},
	Description: "example",
	Cooldown:    2 * time.Second,
	Run: func(s *discordgo.Session, m *discordgo.MessageCreate, ctx []string) {

		s.ChannelMessageSend(m.ChannelID, "example")

	},
}

func init() {
	handlers.LoadCommands(example)
}
