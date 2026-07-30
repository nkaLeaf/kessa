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
	if m.Author.ID != "283723843216867328" {
		return
	}

	args, nan := strings.CutPrefix(m.Content, handlers.Cfg.Prefix)
	msg := strings.Fields(args)
	msg[0] = strings.ToLower(msg[0])
	handlers.Info(strings.Join(msg[1:], " "))

	switch nan {
	case true:
		cmd, ok := handlers.CommandMap[msg[0]]
		if ok {
		} else if !ok {
			if cmd, ok = handlers.CommandMap[handlers.AliasesMap[msg[0]]]; ok {
			} else {
				return
			}
		}
		if cmd != nil {
			cmd.Run(s, m, msg[1:])
			handlers.Logger(cmd.Name, logrus.DebugLevel, errors.New(" named command ran"))
		}
	case false:
		return
	}

}
func init() {
	handlers.EventMap["messageCreate"] = messageCreate
}
