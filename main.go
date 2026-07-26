package main

import (
	_ "kessa/commands"
	_ "kessa/events"
	"kessa/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {

	dg, err := discordgo.New("Bot " + handlers.Cfg.Token)
	handlers.Logger("session create", logrus.FatalLevel, err)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent

	handlers.LoadEvents(dg)
	handlers.LoadCommands()

	err = dg.Open()
	handlers.Info("session open", logrus.InfoLevel)
	handlers.Logger("session couldnt open", logrus.FatalLevel, err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
