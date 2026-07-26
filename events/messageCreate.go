package events

import (
	"errors"
	"kessa/handlers"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	msg, nan := strings.CutPrefix(m.Content, handlers.Cfg.Prefix)
	switch nan {
	case true:
		if cmd, ok := handlers.CommandMap[msg]; ok {
			cmd.Run(s, m)
			handlers.Logger(cmd.Name+" named command ran", logrus.DebugLevel, errors.New("duh"))
		}
	case false:
		return
	}

}
func init() {
	handlers.EventMap["messageCreate"] = messageCreate
}
