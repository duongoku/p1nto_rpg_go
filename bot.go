package p1nto

import (
	"fmt"
	"strings"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var (
	mess = make(chan *discordgo.MessageCreate, 1)
	prefix = "p."
	terminate = make(chan bool, 1)
	users = make(map[string]player)
)

func RNG(x int) bool {
	//x% is the miss chance
	if rand.Intn(100) < x {
		return false
	} else {
		return true
	}
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	mess<- m
}

func MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := m.Content

	if !strings.HasPrefix(content, prefix) {
		return
	}

	content = content[len(prefix):]
	tmp := true

	switch(tmp){
		case strings.HasPrefix(content, "die"):
			terminate<- true
			break
		case strings.HasPrefix(content, "combat"):
			CombatHandle(s, m)
			break
		case strings.HasPrefix(content, "help"):
			HelpHandle(s, m)
			break
		default:
			break
	}
	fmt.Println(m.Author.ID + ": " + content)
}

func Loop(s *discordgo.Session, stopListening func()) {
	for {
		select {
			case m := <-mess:
				MessageHandle(s, m)
			case <-terminate:
				fmt.Println("Terminated!")
				return
		}
	}
}