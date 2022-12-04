package command

import (
	"log"

	"github.com/18lab/discord-bot/client/config"
	"github.com/bwmarrin/discordgo"
)

var (
	List                         = []*discordgo.ApplicationCommand{}
	Private                      = []*discordgo.ApplicationCommand{}
	permissionManageServer int64 = discordgo.PermissionManageServer
)

// DeployProd deploys all commands into production.
func DeployProduction(filename string) {
	config := config.Load(filename)

	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredCommands, err := s.ApplicationCommandBulkOverwrite(config.AppID, "", List)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash public commands successfully to production.", len(registeredCommands))

	for _, pc := range Private {
		cmd, err := s.ApplicationCommandCreate(config.AppID, config.GuildID, pc)
		if err != nil {
			log.Fatalln("Cannot create commands:", err)
		}
		log.Printf("Deploy private command '%v' to production.", cmd.Name)
	}
}

// DeployProd deploys all commands into development.
func DeployTest() {
	config := config.Load("config.json")

	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln(`Error setting up discord client:`, err)
	}

	registeredCommands, err := s.ApplicationCommandBulkOverwrite(config.AppID, config.GuildID, List)
	if err != nil {
		log.Fatalln("Cannot create commands:", err)
	}
	log.Printf("Deployed %v slash public commands successfully to development.", len(registeredCommands))

	for _, pc := range Private {
		cmd, err := s.ApplicationCommandCreate(config.AppID, config.GuildID, pc)
		if err != nil {
			log.Fatalln("Cannot create commands:", err)
		}
		log.Printf("Deploy private command '%v' to development.", cmd.Name)
	}
}
