package p1nto

import (
	//"fmt"
	//"strings"

	"github.com/bwmarrin/discordgo"
)

func HelpHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	temp := "```Declare prefix 'p.' to send message into bot.\nAfter prefix stating a command to play the game\n"
	temp += "- 'combat + @playername' to duel with @playername\n"
	temp += "- 'stats' to display your stats\n"
	temp += "- 'equipment' to display your equipment\n"
	temp += "- 'shop' to go to shop and see shelves\n"
	temp += "- 'help' to display this shit```\n"

	m.Content = temp
	s.ChannelMessageSend(m.ChannelID, m.Content)
}
