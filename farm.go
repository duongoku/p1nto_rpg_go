package p1nto

import (
	"strings"
	"strconv"
	// "fmt"

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
	NPCs[0] = &NPC{"Raptor", []int{0}, 10, 40, 8, 3, 1, 50}
	NPCs[1] = &NPC{"Murk Wolf", []int{0}, 20, 50, 25, 5, 1, 1}
	NPCs[2] = &NPC{"Krug", []int{0}, 25, 60, 15, 5, 1, 50}
	NPCs[3] = &NPC{"Gromp", []int{0}, 35, 100, 30, 10, 5, 1}
	NPCs[4] = &NPC{"Sentinel", []int{0}, 60, 100, 50, 12, 10, 10}
	NPCs[5] = &NPC{"Brambleback", []int{0}, 100, 180, 50, 25, 15, 10}
	NPCs[6] = &NPC{"Rift Scuttler", []int{0}, 150, 280, 30, 25, 20, 60}
	NPCs[7] = &NPC{"Earth Dragon", []int{0}, 170, 300, 40, 30, 5, 15}
	NPCs[8] = &NPC{"Fire Dragon", []int{0}, 210, 300, 70, 20, 22, 25}
	NPCs[9] = &NPC{"Green Hell Dragon", []int{0}, 300, 450, 100, 40, 20, 29}
	NPCs[10] = &NPC{"Omega Shenron", []int{0}, 420, 666, 150, 50, 33, 35}
	NPCs[11] = &NPC{"Baron Nashor", []int{0}, 600, 700, 220, 70, 15, 33}
	NPCs[12] = &NPC{"Roshan", []int{0}, 750, 888, 350, 100, 40, 33}
	NPCs[13] = &NPC{"Lich King", []int{0}, 1250, 1000, 600, 150, 35, 40}
	NPCs[14] = &NPC{"The Ultimate P1nto Conqueror", []int{0}, 3000, 2500, 1000, 333, 50, 50}
}

func InitDungeon() {
	dungeons[0] = &dungeon{[]int{0, 1, 2, 3}, []int{50, 30, 15, 5}}
	dungeons[1] = &dungeon{[]int{4, 5, 6}, []int{60, 35, 5}}
	dungeons[2] = &dungeon{[]int{7, 8, 9, 10}, []int{40, 30, 20, 10}}
	dungeons[3] = &dungeon{[]int{11, 12, 13}, []int{50, 30, 20}}
	dungeons[4] = &dungeon{[]int{14}, []int{100}}
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
		tmp += n.Name + "'s HP reached 0, " + p.Name + " defeated " + n.Name
		tmp += "and claimed $" + strconv.Itoa(n.Bounty)
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
		tmp += p.Name + "'s HP reached 0, " + p.Name + " is defeated"
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

	CombatNPC(s, m, players[m.Author.ID], NPCs[dun.Monsters[cnt]])
}