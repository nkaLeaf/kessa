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
			handlers.Logger(cmd.Name, logrus.DebugLevel, errors.New(" named command ran"))
			cmd.Run(s, m)
		}
	case false:
		return
	}

}
func init() {
	handlers.EventMap["messageCreate"] = messageCreate
}
