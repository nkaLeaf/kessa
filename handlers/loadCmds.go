package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Cooldown    time.Duration
	Run         func(*discordgo.Session, *discordgo.MessageCreate)
}

var CommandMap = make(map[string]*Command)

func LoadCommands() {
	for name := range CommandMap {
		Info("Command: "+name+" has loaded", logrus.InfoLevel)
	}
}
