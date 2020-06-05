package p1nto

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func InventoryHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	content := strings.Split(m.Content, " ")
	if len(content)>2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	if len(m.Mentions)>1 {
		s.ChannelMessageSend(m.ChannelID, "You mentioned more than one person")
		return
	}
	u := m.Author
	if len(m.Mentions)==1 {
		u = m.Mentions[0]
	}

	CheckPlayer(u)
	p := users[u.ID]

	tmps := (p.Name + "'s Inventory:\n")
	tmps += "```Value type in order: Hp, Price, Atk, Def, Evasion, CritChance\n"
	for i, inv := range p.Inventory {
		tmps += "Slot " + strconv.Itoa(i) + ": "
		tmps += inv.Name + ": " + "\n"
		tmps += strconv.Itoa(inv.Price) + " "
		tmps += strconv.Itoa(inv.Hp) + " "
		tmps += strconv.Itoa(inv.Atk) + " "
		tmps += strconv.Itoa(inv.Def) + " "
		tmps += strconv.Itoa(inv.Evasion) + " "
		tmps += strconv.Itoa(inv.CritChance) + "\n"
	}
	tmps += "```"
	s.ChannelMessageSend(m.ChannelID, tmps)
}