package commands

import (
	"kessa/handlers"
	"time"

	"github.com/bwmarrin/discordgo"
)

func run(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, "WebSocket Ping: "+s.HeartbeatLatency().Round(time.Millisecond).String()+"ms")

}
func init() {
	handlers.CommandMap["ping"] = &handlers.Command{
		Name:        "ping",
		Aliases:     []string{"p"},
		Description: "Check the bot's latency",
		Cooldown:    2 * time.Second,
		Run:         run,
	}
}
