package p1nto

import (
	"github.com/bwmarrin/discordgo"
)

type item struct {
	Name string
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

func Equip(user *discordgo.User, itemID int) {
	users[user.ID].Hp += items[itemID].Hp
	users[user.ID].Atk += items[itemID].Atk
	users[user.ID].Def += items[itemID].Def
	users[user.ID].Evasion += items[itemID].Evasion
	users[user.ID].CritChance += items[itemID].CritChance
}