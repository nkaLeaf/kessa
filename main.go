package main

import (
	"fmt"
	_ "kessa/events"
	"kessa/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := handlers.LoadConf()
	handlers.Catch("config load", "fatal", err)

	dg, err := discordgo.New("Bot " + cfg.Token)
	handlers.Catch("session create", "fatal", err)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent

	handlers.LoadEvents(dg)

	err = dg.Open()
	handlers.Catch("session open", "fatal", err)

	fmt.Println("bot is running now good job bitchass the prefix is: ", cfg.Prefix)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
