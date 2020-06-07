package p1nto

import (
	"strings"
	"strconv"
	// "time"
	// "fmt"

	"github.com/bwmarrin/discordgo"
)

func DamageCalc(hp *int, atk int, def int, crt int, eva int) (bool, bool, int) {
	a, b := false, false
	if RNG(eva) {
		atk = 0
		a = true
	}
	if RNG(crt) {
		atk += atk
		b = true
	}
	atk = Max(0, atk - def)
	*hp = *hp - atk
	return a, b, atk
}

func Hit(s *discordgo.Session, m **discordgo.Message, chanID *string, p1 *player, p2 *player, hp1 *int, hp2 *int) bool {
	*m, _ = s.ChannelMessage(*chanID, (*m).ID)
	a, b, c := DamageCalc(hp2, p1.Atk, p2.Def, p1.CritChance, p2.Evasion)
	tmps := ((*m).Content + "\n")
	if a {
		tmps += p1.Name + " missed!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	} else if b {
		tmps += p1.Name + " landed a critical hit on " + p2.Name + " for "
		tmps += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	} else {
		tmps += p1.Name + " hit " + p2.Name + " for "
		tmps += strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	}
	if *hp2 <= 0 {
		*m, _ = s.ChannelMessage(*chanID, (*m).ID)
		tmps = (*m).Content + "\n"
		tmps += p2.Name + "'s HP reached 0, " + p2.Name + " is dead.\n"
		tmps += p1.Name + " Won! $10 has been added to " + p1.Name + "'s balance"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
		p1.Money += 10
		return true
	}
	return false
}

func CombatHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")

	if len(content)<2 {
		s.ChannelMessageSend(m.ChannelID, "Need one argument")
		return
	}
	if len(content)>2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	if len(m.Mentions)<1 {
		s.ChannelMessageSend(m.ChannelID, "You must mention a person")
		return
	}
	if len(m.Mentions)>1 {
		s.ChannelMessageSend(m.ChannelID, "You mentioned more than one person")
		return
	}
	if m.Mentions[0].ID == m.Author.ID {
		s.ChannelMessageSend(m.ChannelID, "You can't kick your own butt bro")
		return
	}

	u1 := m.Author
	u2 := m.Mentions[0]

	CheckPlayer(u1)
	CheckPlayer(u2)

	p1 := players[u1.ID]
	p2 := players[u2.ID]

	hp1 := p1.Hp
	hp2 := p2.Hp

	//u1 goes first since u1 declares the battle

	tmpmm, _ := s.ChannelMessageSend(m.ChannelID, p1.Name + " vs " + p2.Name)
	tmpm := &tmpmm

	for {
		if Hit(s, tmpm, &(m.ChannelID), p1, p2, &hp1, &hp2) {
			break
		}
		if Hit(s, tmpm, &(m.ChannelID), p2, p1, &hp2, &hp1) {
			break
		}
		// time.Sleep(500 * time.Millisecond)
	}
}