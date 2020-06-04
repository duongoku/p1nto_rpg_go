package p1nto

import (
	"github.com/bwmarrin/discordgo"
)


type player struct {
	Name string
	Money int
	Equipment []item
	Inventory []item
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

func CheckPlayer(user *discordgo.User) {
	_, ok := users[user.ID]
	if !ok {
		users[user.ID] = &player{user.Username, 0, nil, nil, 50, 10, 1, 1, 1}
		Equip(user, 1)
		Equip(user, 2)
		Equip(user, 3)
	}
	return
}