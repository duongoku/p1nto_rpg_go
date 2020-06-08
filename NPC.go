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

var NPCEmoji = []string{"<:raptor:719483403208228934> ", "<:murkwolf:719483393322254407>", "<:krug:719483382790226012>", "<:gromp:719483369737551902>"}

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

func NPCHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide a NPC ID")
		return
	}
	if len(content) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	NPCID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong format")
		return
	}
	if (NPCID < 0) || (NPCID >= len(NPCs)) {
		s.ChannelMessageSend(m.ChannelID, "NPC not found")
		return
	}

	tmp := "**" + NPCs[NPCID].Name + "** "
	if NPCID < 4 {
		tmp += NPCEmoji[NPCID]
	}

	tmp +=  "\n Loot: "
	tmp +=  "\n Bounty: $" + strconv.Itoa(NPCs[NPCID].Bounty)
	tmp +=  "\n Health Points: " + strconv.Itoa(NPCs[NPCID].Hp) + " "
	tmp +=  "\n Attack: " + strconv.Itoa(NPCs[NPCID].Atk) + " "
	tmp +=  "\n Defense: " + strconv.Itoa(NPCs[NPCID].Def) + " "
	tmp +=  "\n Evasion: " + strconv.Itoa(NPCs[NPCID].Evasion) + "%"
	tmp +=  "\n Crit Chance: " + strconv.Itoa(NPCs[NPCID].CritChance) + "%"

	s.ChannelMessageSend(m.ChannelID, tmp)
}