package p1nto

import (
	"strconv"
	"strings"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type item struct {
	Name string
	Price int
	SlotID int
	Hp int
	Atk int
	Def int
	Evasion int
	CritChance int
}

func InitItem() {
	//Name, Price, SlotID, Hp, Atk, Def, Evasion, CritChane
	items[0] = &item{"Wooden Amulet", 50, 0, 5, 5, 1, 5, 5}
    items[1] = &item{"Wooden Stick", 25, 1, 0, 10, 0, 0, 20}
    items[2] = &item{"Wooden Plate Mail", 25, 2, 25, 0, 2, 0, 0}
    items[3] = &item{"Wooden Greaves", 25, 3, 5, 0, 1, 20, 0}
}

func EquipHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content)>2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	if len(content)<2 {
		s.ChannelMessageSend(m.ChannelID, "Need one argument")
		return
	}
	InvenID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong format")
		return
	}
	if(InvenID >= len(players[m.Author.ID].Inventory)) {
		s.ChannelMessageSend(m.ChannelID, "You don't have that many items bro")
		return
	}
	if(InvenID < 0) {
		s.ChannelMessageSend(m.ChannelID, "Bro?")
		return
	}

	temp := players[m.Author.ID].Inventory[InvenID]
	if Equip(m.Author, players[m.Author.ID].Inventory[InvenID]) {
		players[m.Author.ID].Inventory = append(players[m.Author.ID].Inventory[:InvenID], players[m.Author.ID].Inventory[InvenID+1:]...)

		s.ChannelMessageSend(m.ChannelID, "Equipped " + items[temp].Name)
	} else {
		s.ChannelMessageSend(m.ChannelID, "That slot is already equipped with another item")
		return
	}
}

func UnequipHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content)>2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}
	if len(content)<2 {
		s.ChannelMessageSend(m.ChannelID, "Need one argument")
		return
	}
	SlotID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong format")
		return
	}
	if(SlotID >= len(players[m.Author.ID].Equipment)) {
		s.ChannelMessageSend(m.ChannelID, "You don't have that many slots bro")
		return
	}
	
	temp := players[m.Author.ID].Equipment[SlotID]
	if Unequip(m.Author, SlotID) {
		players[m.Author.ID].Inventory = append(players[m.Author.ID].Inventory, temp)

		s.ChannelMessageSend(m.ChannelID, "Unequipped " + items[temp].Name)
	} else {
		s.ChannelMessageSend(m.ChannelID, "That slot is not equipped with anything")
		return
	}
}

func Equip(user *discordgo.User, itemID int) bool {
	if(players[user.ID].Equipment[items[itemID].SlotID] == -1) {
		players[user.ID].Equipment[items[itemID].SlotID] = itemID
		temp := items[itemID]

		players[user.ID].Hp += temp.Hp
		players[user.ID].Atk += temp.Atk
		players[user.ID].Def += temp.Def
		players[user.ID].Evasion += temp.Evasion
		players[user.ID].CritChance += temp.CritChance

		fmt.Println("Equipped " + temp.Name + " for " + players[user.ID].Name)
		return true
	} else {
		return false
	}
}

func Unequip(user *discordgo.User, SlotID int) bool {
	if(players[user.ID].Equipment[SlotID] != -1) {
		temp := items[players[user.ID].Equipment[SlotID]]

		players[user.ID].Hp -= temp.Hp
		players[user.ID].Atk -= temp.Atk
		players[user.ID].Def -= temp.Def
		players[user.ID].Evasion -= temp.Evasion
		players[user.ID].CritChance -= temp.CritChance

		players[user.ID].Equipment[SlotID] = -1

		fmt.Println("Unequipped " + temp.Name + " for " + players[user.ID].Name)
		return true
	} else {
		return false
	}
}