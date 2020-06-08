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

func InitNPC() {
	//Credited to github.com/Darthveloper21

	//Name, Loot[], Bounty, Hp, Atk, Def,  Evasion, CritChance
	NPCs[0] = &NPC{"Raptor", []int{0}, 2, 50, 10, 1, 1, 1}
	NPCs[1] = &NPC{"Murk Wolf", []int{0}, 5, 75, 15, 5, 1, 10}
	NPCs[2] = &NPC{"Krug", []int{0}, 10, 100, 20, 10, 5, 50}
	NPCs[3] = &NPC{"Gromp", []int{0}, 20, 150, 35, 20, 5, 35}
	NPCs[4] = &NPC{"Sentinel", []int{0}, 25, 240, 65, 45, 10, 10}
	NPCs[5] = &NPC{"Brambleback", []int{0}, 35, 340, 80, 60, 15, 15}
	NPCs[6] = &NPC{"Rift Scuttler", []int{0}, 50, 500, 55, 100, 20, 90}
	NPCs[7] = &NPC{"Earth Dragon", []int{0}, 70, 650, 100, 150, 25, 15}
	NPCs[8] = &NPC{"Fire Dragon", []int{0}, 85, 550, 150, 100, 15, 30}
	NPCs[9] = &NPC{"Green Hell Dragon", []int{0}, 100, 700, 200, 185, 20, 29}
	NPCs[10] = &NPC{"Omega Shenron", []int{0}, 150, 1000, 300, 230, 33, 40}
	NPCs[11] = &NPC{"Baron Nashor", []int{0}, 200, 1500, 400, 300, 20, 33}
	NPCs[12] = &NPC{"Roshan", []int{0}, 225, 2000, 500, 350, 33, 33}
	NPCs[13] = &NPC{"Lich King", []int{0}, 250, 2500, 750, 500, 35, 40}
	NPCs[14] = &NPC{"The Ultimate P1nto Conqueror", []int{0}, 1000, 3500, 1200, 600, 50, 50}
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

	tmp := "**" + NPCs[NPCID].Name + "**"
	tmp +=  "\n Loot: "
	tmp +=  "\n Bounty: $" + strconv.Itoa(NPCs[NPCID].Bounty)
	tmp +=  "\n Health Points: " + strconv.Itoa(NPCs[NPCID].Hp) + " "
	tmp +=  "\n Attack: " + strconv.Itoa(NPCs[NPCID].Atk) + " "
	tmp +=  "\n Defense: " + strconv.Itoa(NPCs[NPCID].Def) + " "
	tmp +=  "\n Evasion: " + strconv.Itoa(NPCs[NPCID].Evasion) + "%"
	tmp +=  "\n Crit Chance: " + strconv.Itoa(NPCs[NPCID].CritChance) + "%"

	s.ChannelMessageSend(m.ChannelID, tmp)
}