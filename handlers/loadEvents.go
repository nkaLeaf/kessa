package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var EventMap = make(map[string]interface{})

func LoadEvents(dg *discordgo.Session) {
	for name, event := range EventMap {
		fmt.Println("event " + name + " has loaded")
		dg.AddHandler(event)
	}

}
