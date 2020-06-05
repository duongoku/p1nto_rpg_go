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
	var token string = os.Getenv("TOKEN_test")
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
	
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is online!")

	stopListening := dg.AddHandler(p1nto.MessageCreate)

	p1nto.Loop(dg, stopListening)
}