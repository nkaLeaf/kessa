package kessa

import (
	"bufio"
	_ "kessa/commands"
	_ "kessa/events"
	"kessa/handlers"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {

	dg, err := discordgo.New("Bot " + handlers.Cfg.Token)
	handlers.Logger("session create", logrus.FatalLevel, err)

	dg.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent

	handlers.LoadEvents(dg)
	handlers.OpenDB()

	err = dg.Open()
	handlers.Info("session open")
	handlers.Logger("session couldnt open", logrus.FatalLevel, err)

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		handlers.Logger("couldnt read command line", logrus.WarnLevel, err)

		input = strings.TrimSpace(input)

		switch input {
		case "exit":
			handlers.Info("exiting bot")
			dg.Close()
			os.Exit(0)
		default:
			handlers.Info("unknown code broski")
			continue
		}
	}
}
