package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Cooldown    time.Duration
	Run         func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

var CommandMap = make(map[string]*Command)
var AliasesMap = make(map[string]string)

func LoadCommands(cmd *Command) {
	CommandMap[cmd.Name] = cmd
	if cmd.Aliases == nil {
		return
	}
	for _, alias := range cmd.Aliases {
		AliasesMap[alias] = cmd.Name
	}
	Info("Command: " + cmd.Name + " has loaded")

}
