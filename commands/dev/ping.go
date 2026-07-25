package commands

import (
	"kessa/handlers"
	"time"

	"github.com/bwmarrin/discordgo"
)

func run(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, "WebSocket Ping:  ms "+s.HeartbeatLatency().Round(time.Millisecond).String())

}
func init() {
	handlers.CommandMap["ping"] = &handlers.Command{
		Name:        "ping",
		Aliases:     []string{"p"},
		Description: "",
		Cooldown:    2 * time.Second,
		Run:         run,
	}
}
