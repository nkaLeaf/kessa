package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var EventMap = make(map[string]interface{})

func LoadEvents(dg *discordgo.Session) {
	for name, event := range EventMap {
		Info("Event: " + name + " has loaded")
		dg.AddHandler(event)
	}

}
