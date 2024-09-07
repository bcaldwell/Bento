package responsehelpers

import (
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func SendMessageFromList(s *discordgo.Session, reference *discordgo.MessageReference, options []string) {
	s.ChannelMessageSendReply(reference.ChannelID, options[rand.Int()%len(options)], reference)
}
