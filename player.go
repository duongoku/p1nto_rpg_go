package p1nto

import (
	"github.com/bwmarrin/discordgo"
)


type player struct {
	Name string
	Money int
	Equipment [5]int
	Inventory []int
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

func CheckPlayer(user *discordgo.User) {
	_, ok := players[user.ID]
	if !ok {
		//Basic stats are 50, 10, 1, 0, 0
		players[user.ID] = &player{user.Username, 0, [5]int{-1, -1, -1, -1, -1}, nil, 50, 10, 1, 0, 0}
	}
	return
}