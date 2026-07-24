package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "go is a good language fr fr")
	fmt.Println("im up bro")
}
