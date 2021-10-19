package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/duongoku/p1nto"
)

func main() {
	var token string = os.Getenv("TOKEN")
	rand.Seed(time.Now().UnixNano())
	if token == "" {
		fmt.Println("No TOKEN found!")
		return
	}
	
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	err = dg.Open()
	defer dg.Close()

	dg.UpdateStatus(0, "p.help")
	
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is online!")

	dg.AddHandler(p1nto.MessageCreate)
	p1nto.Loop(dg)
}