package p1nto

import (
	"strings"
	"strconv"
	"fmt"
	// "math/rand"

	"github.com/bwmarrin/discordgo"
)

type NPC struct {
	Name string
	Loot []int
	Bounty int
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

type dungeon struct {
	Monsters []int
	SpawnRate []int
}

func InitNPC() {
	NPCs[0] = &NPC{"Raptor", []int{0}, 10, 100, 30, 5, 1, 1}
	NPCs[1] = &NPC{"Murk Wolf", []int{0}, 50, 200, 60, 8, 1, 1}
	NPCs[2] = &NPC{"Krug", []int{0}, 100, 300, 80, 10, 1, 1}
	NPCs[3] = &NPC{"Gromp", []int{0}, 200, 400, 100, 12, 1, 1}
}

func InitDungeon() {
	dungeons[0] = &dungeon{[]int{0, 1, 2, 3}, []int{50, 30, 15, 5}}
}

func HitNPC(s *discordgo.Session, m *discordgo.Message, p *player, n *NPC, hpp *int, hpn *int) bool {
	m, _ = s.ChannelMessage(m.ChannelID, m.ID)
	tmp := m.Content + "\n"
	
	a, b, c := DamageCalc(hpn, p.Atk, n.Def, p.CritChance, n.Evasion)
	if a {
		tmp += p.Name + " missed!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else if b {
		tmp += p.Name + " landed a critical hit on " + n.Name + " for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else {
		tmp += p.Name + " hit " + n.Name + " for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	}
	if *hpn <= 0 {
		m, _ = s.ChannelMessage(m.ChannelID, m.ID)
		tmp = m.Content + "\n"
		tmp += n.Name + "'s HP reached 0. " + p.Name + " defeated " + n.Name
		tmp += ", " + p.Name + " got $" + strconv.Itoa(n.Bounty)
		p.Money += n.Bounty
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
		return true
	}

	m, _ = s.ChannelMessage(m.ChannelID, m.ID)
	tmp = m.Content + "\n"
	
	a, b, c = DamageCalc(hpp, n.Atk, p.Def, n.CritChance, p.Evasion)
	if a {
		tmp += n.Name + " missed!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else if b {
		tmp += n.Name + " landed a critical hit on " + p.Name + " for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	} else {
		tmp += n.Name + " hit " + p.Name + " for "
		tmp += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
	}
	if *hpp <= 0 {
		m, _ = s.ChannelMessage(m.ChannelID, m.ID)
		tmp = m.Content + "\n"
		tmp += p.Name + "'s HP reached 0. " + p.Name + " is defeated"
		s.ChannelMessageEdit(m.ChannelID, m.ID, tmp)
		return true
	}

	return false
}

func CombatNPC(s *discordgo.Session, m *discordgo.MessageCreate, p *player, n *NPC) {
	tmpm, _ := s.ChannelMessageSend(m.ChannelID, p.Name + " found a " + n.Name)

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

	fmt.Println("Found " + NPCs[dun.Monsters[cnt]].Name)

	CombatNPC(s, m, players[m.Author.ID], NPCs[dun.Monsters[cnt]])
}