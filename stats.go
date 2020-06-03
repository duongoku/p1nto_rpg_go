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
	p := users[u.ID]

	/*Name string
	Money int
	Equipment []int
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
	tmps = tmps + "\nEvasion: " + strconv.Itoa(p.Evasion)
	tmps = tmps + "\nCritical Chance: " + strconv.Itoa(p.CritChance)

	s.ChannelMessageSend(m.ChannelID, tmps)
}