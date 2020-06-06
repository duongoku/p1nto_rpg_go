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

	tmps := "```" + p.Name + "'s Inventory:\n"
	for invID, eqm := range p.Inventory {
		tmps = tmps + "\n" + "Inventory ID: " + strconv.Itoa(invID) + "\n"
		tmps += "Type:" + strconv.Itoa(items[eqm].SlotID) + " " + items[eqm].Name + ": "
		if items[eqm].Hp > 0 {
			tmps = tmps + "Hp+" + strconv.Itoa(items[eqm].Hp) + ", "
		}
		if items[eqm].Atk > 0 {
			tmps = tmps + "Attack+" + strconv.Itoa(items[eqm].Atk) + ", "
		}
		if items[eqm].Def > 0 {
			tmps = tmps + "Defense+" + strconv.Itoa(items[eqm].Def) + ", "
		}
		if items[eqm].Evasion > 0 {
			tmps = tmps + "Evasion+" + strconv.Itoa(items[eqm].Evasion) + "%, "
		}
		if items[eqm].CritChance > 0 {
			tmps = tmps + "Critical Chance+" + strconv.Itoa(items[eqm].CritChance) + "%, "
		}
		tmps = tmps[0:len(tmps)-2]
	}
	tmps += "```"
	s.ChannelMessageSend(m.ChannelID, tmps)
}