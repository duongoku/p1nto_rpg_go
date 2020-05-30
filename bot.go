package p1nto

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	mess = make(chan *discordgo.MessageCreate, 1)
	prefix = "p."
	terminate = make(chan bool, 1)
)

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

	switch(content){
		case "die":
			terminate<- true
			break
		default:
			break
	}

	fmt.Println(content)
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