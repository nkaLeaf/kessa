package commands

import (
	"fmt"
	"kessa/handlers"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var eval = &handlers.Command{
	Name:        "eval",
	Aliases:     []string{""},
	Description: "eval",
	Cooldown:    2 * time.Second,
	Run: func(s *discordgo.Session, m *discordgo.MessageCreate, ctx []string) {
		if len(ctx) == 0 {
			s.ChannelMessageSend(m.ChannelID, "there are no args")
			return
		} else {
			i := interp.New(interp.Options{})
			i.Use(stdlib.Symbols)
			i.Eval(`import (
					"fmt"
					"kessa/handlers"
					"strings"
					"time")
					`)
			r, err := i.Eval(strings.Join(ctx, ""))
			handlers.Logger("eval err", logrus.WarnLevel, err)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "youre cooking shit go back to js bro")
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprint(r.Interface()))
			}
		}

	},
}

func init() {
	handlers.LoadCommands(eval)
}
