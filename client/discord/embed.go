package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	ISO8601        = "2006-01-02T03:04:05-0700"
	DateTimeFormat = "Jan 02 15:04"
	EmbedTimestamp = time.Now().Format(ISO8601)
	EmbedColor     = 15548997
	EmbedFooter    = func(s *discordgo.Session) *discordgo.MessageEmbedFooter {
		return &discordgo.MessageEmbedFooter{
			Text:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(""),
		}
	}
)

// NewEmbed creates a new MessageEmbed with some initial templated fields, and returns its address.
func NewEmbed(s *discordgo.Session) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:     EmbedColor,
		Timestamp: EmbedTimestamp,
		Footer:    EmbedFooter(s),
	}
}
