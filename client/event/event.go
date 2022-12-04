package event

import (
	"log"
	"os"

	"github.com/18lab/discord-bot/schema"
	"github.com/bwmarrin/discordgo"
)

var (
	List = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

func init() {
	// Mkdir all neccessary path
	if err := os.MkdirAll(schema.PathData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
