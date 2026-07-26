package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var EventMap = make(map[string]interface{})

func LoadEvents(dg *discordgo.Session) {
	for name, event := range EventMap {
		Info("Event: "+name+" has loaded", logrus.InfoLevel)
		dg.AddHandler(event)
	}

}
