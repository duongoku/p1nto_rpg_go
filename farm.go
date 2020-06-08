package p1nto

import (
	"strings"
	"strconv"
	// "fmt"

	"github.com/bwmarrin/discordgo"
)

type dungeon struct {
	Monsters []int
	SpawnRate []int
}

func InitDungeon() {
	dungeons[0] = &dungeon{[]int{0, 1, 2, 3}, []int{50, 30, 15, 5}}
	dungeons[1] = &dungeon{[]int{4, 5, 6}, []int{60, 35, 5}}
	dungeons[2] = &dungeon{[]int{7, 8, 9, 10}, []int{40, 30, 20, 10}}
	dungeons[3] = &dungeon{[]int{11, 12, 13}, []int{60, 30, 10}}
	dungeons[4] = &dungeon{[]int{14}, []int{100}}
}

func HitNPC(s *discordgo.Session, m *discordgo.Message, p *player, n *NPC, hpp *int, hpn *int) bool {
	m, _ = s.ChannelMessage(m.ChannelID, m.ID)
	tmp := m.Content + "\n"
	
	a, b, c := DamageCalc(hpn, p.Atk, n.Def, p.CritChance, n.Evasion)
	if a {
		tmp += "**" + p.Name + "** missed!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else if b {
		tmp += "**" + p.Name + "** landed a critical hit on **" + n.Name + "** for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else {
		tmp += "**" + p.Name + "** hit **" + n.Name + "** for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	}
	if *hpn <= 0 {
		m, _ = s.ChannelMessage(m.ChannelID, m.ID)
		tmp = m.Content + "\n"
		tmp += "**" + n.Name + "**'s HP reached 0, **" + p.Name + "** defeated **" + n.Name
		tmp += "** and claimed $" + strconv.Itoa(n.Bounty)
		p.Money += n.Bounty
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
		return true
	}

	m, _ = s.ChannelMessage(m.ChannelID, m.ID)
	tmp = m.Content + "\n"
	
	a, b, c = DamageCalc(hpp, n.Atk, p.Def, n.CritChance, p.Evasion)
	if a {
		tmp += "**" + n.Name + "** missed!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else if b {
		tmp += "**" + n.Name + "** landed a critical hit on **" + p.Name + "** for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else {
		tmp += "**" + n.Name + "** hit **" + p.Name + "** for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	}
	if *hpp <= 0 {
		m, _ = s.ChannelMessage(m.ChannelID, m.ID)
		tmp = m.Content + "\n"
		tmp += "**" + p.Name + "**'s HP reached 0, **" + p.Name + "** is defeated"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
		return true
	}

	return false
}

func CombatNPC(s *discordgo.Session, m *discordgo.MessageCreate, p *player, n *NPC) {
	tmpm, _ := s.ChannelMessageSend(m.ChannelID, "A wild **" + n.Name + "** appeared")

	hpp := p.Hp
	hpn := n.Hp

	for {
		if HitNPC(s, tmpm, p, n, &hpp, &hpn) {
			break
		}
	}
}

func FarmHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide a dungeon(stage) ID")
		return
	}
	if len(content) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	dunID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong format")
		return
	}
	if (dunID < 0) || (dunID >= len(dungeons)) {
		s.ChannelMessageSend(m.ChannelID, "Dungeon not found")
		return
	}

	dun := dungeons[dunID]
	var tmp, cnt int = xRNG(), 0

	for tmp >= dun.SpawnRate[cnt] {
		tmp -= dun.SpawnRate[cnt]
		cnt += 1
	}

	//fmt.Println("debug")
	CheckPlayer(m.Author)
	CombatNPC(s, m, players[m.Author.ID], NPCs[dun.Monsters[cnt]])
}

func DungeonHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide a Dungeon ID")
		return
	}
	if len(content) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	dunID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong format")
		return
	}
	if (dunID < 0) || (dunID >= len(dungeons)) {
		s.ChannelMessageSend(m.ChannelID, "Dungeon not found")
		return
	}

	tmp := "Dungeon " + strconv.Itoa(dunID) + " Details:"
	dun := dungeons[dunID]
	for _, i := range dun.Monsters {
		tmp += "\n**(ID"+ strconv.Itoa(i) + ") " + NPCs[i].Name + ":** "
		tmp +=  strconv.Itoa(NPCs[i].Hp) + " Health Points"
		tmp +=  ", " + strconv.Itoa(NPCs[i].Atk) + " Attack"
		tmp +=  ", " + strconv.Itoa(NPCs[i].Def) + " Defense"
		tmp +=  ", " + strconv.Itoa(NPCs[i].Evasion) + "% Evasion"
		tmp +=  ", " + strconv.Itoa(NPCs[i].CritChance) + "% Crit Chance"
	}

	s.ChannelMessageSend(m.ChannelID, tmp)
}