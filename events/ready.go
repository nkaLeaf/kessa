package events

import (
	"fmt"
	"kessa/handlers"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "go is a good language fr fr")
	fmt.Println("im up bro")
}
func init() {

	handlers.EventMap["ready"] = ready
}
