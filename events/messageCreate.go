package events

import (
	"kessa/handlers"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	msg, nan := strings.CutPrefix(m.Content, handlers.Cfg.Prefix)
	switch nan {
	case true:
		if msg == "kurwa" {
			s.ChannelMessageSend(m.ChannelID, "kurwa")
		}
		if msg == "nacho" {
			s.ChannelMessageSend(m.ChannelID, "bronacho")
		}
	case false:
		return
	}

}
func init() {

	handlers.EventMap["messageCreate"] = messageCreate
}
