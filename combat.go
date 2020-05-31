package p1nto

import (
	// "fmt"
	// "strings"

	"github.com/bwmarrin/discordgo"
)

func CombatHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSend(m.ChannelID, m.Content)
}