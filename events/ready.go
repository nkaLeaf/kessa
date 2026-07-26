package events

import (
	"kessa/handlers"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "go is a good language fr fr")
}
func init() {
	handlers.EventMap["ready"] = ready
}
