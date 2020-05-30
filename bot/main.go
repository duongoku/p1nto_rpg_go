package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/duongoku/p1nto"
)

func main() {
	var token string = os.Getenv("TOKEN")
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
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is online!")
	fmt.Println("Press CTRL+C to exit the program.")

	stopListening := dg.AddHandler(p1nto.MessageCreate)

	p1nto.Loop(dg, stopListening)

	dg.Close()
}