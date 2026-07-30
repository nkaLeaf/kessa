package commands

import (
	"kessa/handlers"
	"time"

	"github.com/bwmarrin/discordgo"
)

var ping = &handlers.Command{
	Name:        "ping",
	Aliases:     []string{"p"},
	Description: "Check the bot's latency",
	Cooldown:    2 * time.Second,
	Run: func(s *discordgo.Session, m *discordgo.MessageCreate, ctx []string) {

		s.ChannelMessageSend(m.ChannelID, "WS Ping: "+s.HeartbeatLatency().Round(time.Millisecond).String())

	},
}

func init() {
	handlers.LoadCommands(ping)
}
