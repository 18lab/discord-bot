package event

import (
	"fmt"
	"log"
	"strings"

	"github.com/18lab/discord-bot/client/discord"
	"github.com/bwmarrin/discordgo"
)

func init() {
	List["stats"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		guildList := []string{}
		for _, g := range s.State.Guilds {
			guildList = append(guildList, g.Name)
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:     fmt.Sprint(s.State.User.String(), " - Stats"),
						Color:     discord.EmbedColor,
						Timestamp: discord.EmbedTimestamp,
						Footer:    discord.EmbedFooter(s),
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:  "In Guilds",
								Value: discord.FieldStyle(len(s.State.Guilds)),
							},
							{
								Name:  "Guild List",
								Value: discord.FieldStyle(strings.Join(guildList, ", ")),
							},
						},
					},
				},
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
