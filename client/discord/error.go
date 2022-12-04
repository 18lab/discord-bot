package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Error Code
const (
	ErrCodeInvalidInteraction = 1001
)

func EmbedError(s *discordgo.Session, i *discordgo.InteractionCreate, code int) {
	embed := NewEmbed(s)
	embed.Title = "Error"

	// Error Code Switch
	switch code {
	case 1001:
		embed.Description = "Invalid interaction. This interaction is intended for original user only."
	}

	embeds := []*discordgo.MessageEmbed{embed}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		ExitErr(err)
	}
}

func ExitErr(err error) {
	// Handle any operations before exiting with an error.

	// Log the error and exit.
	log.Println("Exiting with an error:", err)
	os.Exit(1)
}
