package handlers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
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
		fmt.Println("command " + name + " has loaded")
	}
}
