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
	if(InvenID >= len(users[m.Author.ID].Inventory)) {
		s.ChannelMessageSend(m.ChannelID, "You don't have that many items bro")
		return
	}

	temp := users[m.Author.ID].Inventory[InvenID]
	if Equip(m.Author, users[m.Author.ID].Inventory[InvenID]) {
		users[m.Author.ID].Inventory = append(users[m.Author.ID].Inventory[:InvenID], users[m.Author.ID].Inventory[InvenID+1:]...)

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
	if(SlotID >= len(users[m.Author.ID].Equipment)) {
		s.ChannelMessageSend(m.ChannelID, "You don't have that many slots bro")
		return
	}
	
	temp := users[m.Author.ID].Equipment[SlotID]
	if Unequip(m.Author, SlotID) {
		users[m.Author.ID].Inventory = append(users[m.Author.ID].Inventory, temp)

		s.ChannelMessageSend(m.ChannelID, "Unequipped " + items[temp].Name)
	} else {
		s.ChannelMessageSend(m.ChannelID, "That slot is not equipped with anything")
		return
	}
}

func Equip(user *discordgo.User, itemID int) bool {
	if(users[user.ID].Equipment[items[itemID].SlotID] == -1) {
		users[user.ID].Equipment[items[itemID].SlotID] = itemID
		temp := items[itemID]

		users[user.ID].Hp += temp.Hp
		users[user.ID].Atk += temp.Atk
		users[user.ID].Def += temp.Def
		users[user.ID].Evasion += temp.Evasion
		users[user.ID].CritChance += temp.CritChance

		fmt.Println("Equipped " + temp.Name + " for " + users[user.ID].Name)
		return true
	} else {
		return false
	}
}

func Unequip(user *discordgo.User, SlotID int) bool {
	if(users[user.ID].Equipment[SlotID] != -1) {
		temp := items[users[user.ID].Equipment[SlotID]]

		users[user.ID].Hp -= temp.Hp
		users[user.ID].Atk -= temp.Atk
		users[user.ID].Def -= temp.Def
		users[user.ID].Evasion -= temp.Evasion
		users[user.ID].CritChance -= temp.CritChance

		users[user.ID].Equipment[SlotID] = -1

		fmt.Println("Unequipped " + temp.Name + " for " + users[user.ID].Name)
		return true
	} else {
		return false
	}
}