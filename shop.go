package p1nto

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func ShopHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	temp := "```Value type in order: Hp, Price, Atk, Def, Evasion, CritChance\n"
	for _, value := range items {
		temp += value.Name + ": " + "\n"
		temp += strconv.Itoa(value.Price) + " "
		temp += strconv.Itoa(value.Hp) + " "
		temp += strconv.Itoa(value.Atk) + " "
		temp += strconv.Itoa(value.Def) + " "
		temp += strconv.Itoa(value.Evasion) + " "
		temp += strconv.Itoa(value.CritChance) + "\n"
	}
	temp += "```"
	m.Content = temp
	s.ChannelMessageSend(m.ChannelID, m.Content)
}