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

func Hit(s *discordgo.Session, m **discordgo.Message, chanID *string, u1 *discordgo.User, u2 *discordgo.User, hp1 *int, hp2 *int) bool {
	*m, _ = s.ChannelMessage(*chanID, (*m).ID)
	a, b, c := DamageCalc(hp2, users[u1.ID].Atk, users[u2.ID].Def, users[u1.ID].CritChance, users[u1.ID].Evasion)
	tmps := ((*m).Content + "\n")
	if a {
		tmps = tmps + u1.Username + " missed!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	} else if b {
		tmps = tmps + u1.Username + " landed a critical hit on " + u2.Username + " for "
		tmps = tmps + strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	} else {
		tmps = tmps + u1.Username + " hit " + u2.Username + " for "
		tmps = tmps + strconv.Itoa(c) + " damage!"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
	}
	if *hp2 <= 0 {
		*m, _ = s.ChannelMessage(*chanID, (*m).ID)
		tmps = (*m).Content + "\n"
		tmps = tmps + u2.Username + "'s HP reached 0, " + u2.Username + " is dead.\n"
		tmps = tmps + u1.Username + " Won! $10 has been added to " + u1.Username + "'s balance"
		s.ChannelMessageEdit(*chanID, (*m).ID, tmps)
		users[u1.ID].Money += 10
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

	hp1 := users[u1.ID].Hp
	hp2 := users[u1.ID].Hp

	//u1 goes first since u1 declares the battle

	tmpmm, _ := s.ChannelMessageSend(m.ChannelID, u1.Username + " vs " + u2.Username)
	tmpm := &tmpmm

	for {
		if Hit(s, tmpm, &(m.ChannelID), u1, u2, &hp1, &hp2) {
			break
		}
		if Hit(s, tmpm, &(m.ChannelID), u2, u1, &hp2, &hp1) {
			break
		}
		// time.Sleep(500 * time.Millisecond)
	}
}