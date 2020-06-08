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
		players[user.ID] = &player{user.Username, 0, [5]int{-1, -1, -1, -1, -1}, nil, 50, 10, 1, 1, 1}
		Equip(user, 1)
		Equip(user, 2)
		Equip(user, 3)
	}
	return
}