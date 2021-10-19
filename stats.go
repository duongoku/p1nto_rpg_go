package p1nto

import (
	"strings"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func StatsHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	p := players[u.ID]

	/*Name string
	Money int
	Equipment [5]int
	Inventory []int
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int*/

	tmps := (p.Name + "'s Stats:")
	tmps = tmps + "\nMoney: $" + strconv.Itoa(p.Money)
	tmps = tmps + "\nHealth Points: " + strconv.Itoa(p.Hp)
	tmps = tmps + "\nAttack: " + strconv.Itoa(p.Atk)
	tmps = tmps + "\nDefense: " + strconv.Itoa(p.Def)
	tmps = tmps + "\nEvasion: " + strconv.Itoa(p.Evasion) + "%"
	tmps = tmps + "\nCritical Chance: " + strconv.Itoa(p.CritChance) + "%"

	s.ChannelMessageSend(m.ChannelID, tmps)
}

func EquipmentHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	p := players[u.ID]

	tmps := (p.Name + "'s Equipment:")
	for id, eqmid := range p.Equipment {
		tmps = tmps + "\n" + "Slot " + strconv.Itoa(id) + ": "
		if eqmid==-1 {
			tmps = tmps + "nothing"
			continue
		}
		eqm := items[eqmid]
		tmps = tmps + eqm.Name + ": "
		if eqm.Hp > 0 {
			tmps = tmps + "Hp+" + strconv.Itoa(eqm.Hp) + ", "
		}
		if eqm.Atk > 0 {
			tmps = tmps + "Attack+" + strconv.Itoa(eqm.Atk) + ", "
		}
		if eqm.Def > 0 {
			tmps = tmps + "Defense+" + strconv.Itoa(eqm.Def) + ", "
		}
		if eqm.Evasion > 0 {
			tmps = tmps + "Evasion+" + strconv.Itoa(eqm.Evasion) + "%, "
		}
		if eqm.CritChance > 0 {
			tmps = tmps + "Critical Chance+" + strconv.Itoa(eqm.CritChance) + "%, "
		}
		tmps = tmps[0:len(tmps)-2]
	}

	s.ChannelMessageSend(m.ChannelID, tmps)
}

func InventoryHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	p := players[u.ID]

	tmps := "```" + p.Name + "'s Inventory:"
	if len(p.Inventory) == 0 {
		tmps += "Empty!"
	}

	for invID, eqm := range p.Inventory {
		tmps = tmps + "\n\n" + "Inventory ID: " + strconv.Itoa(invID) + "\n"
		tmps += "Slot ID:" + strconv.Itoa(items[eqm].SlotID) + " " + items[eqm].Name + ": "
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