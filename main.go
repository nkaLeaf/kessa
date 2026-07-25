package main

import (
	_ "kessa/commands"
	_ "kessa/events"
	"kessa/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	dg, err := discordgo.New("Bot " + handlers.Cfg.Token)
	handlers.Catch("session create", "fatal", err)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent

	handlers.LoadEvents(dg)
	handlers.LoadCommands()

	err = dg.Open()
	handlers.Catch("session open", "fatal", err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
