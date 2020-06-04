package p1nto

import (
	"github.com/bwmarrin/discordgo"
)

type item struct {
	Name string
	Price int
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

func InitItem() {
	//Name, Price, Hp, Atk, Def, Evasion, CritChane
	items[0] = item{"Wooden Amulet", 50, 5, 5, 1, 5, 5}
    items[1] = item{"Wooden Stick", 25, 0, 10, 0, 0, 20}
    items[2] = item{"Wooden Plate Mail", 25, 25, 0, 2, 0, 0}
    items[3] = item{"Wooden Greaves", 25, 5, 0, 1, 20, 0}
}

func Equip(user *discordgo.User, itemID int) {
	users[user.ID].Hp += items[itemID].Hp
	users[user.ID].Atk += items[itemID].Atk
	users[user.ID].Def += items[itemID].Def
	users[user.ID].Evasion += items[itemID].Evasion
	users[user.ID].CritChance += items[itemID].CritChance

	users[user.ID].Equipment = append(users[user.ID].Equipment, items[itemID])
}