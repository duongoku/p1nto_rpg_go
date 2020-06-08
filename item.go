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
	items[0] = &item{"Wooden Amulet", 60, 0, 5, 5, 5, 1, 5}
	items[1] = &item{"Wooden Stick", 60, 1, 0, 30, 0, 0, 7}
	items[2] = &item{"Wooden Plate Mail", 60, 2, 50, 0, 10, 0, 0}
	items[3] = &item{"Wooden Greaves", 250, 3, 0, 0, 0, 10, 0}
	items[4] = &item{"Wooden Helmet", 50, 4, 15, 0, 5, 0, 0}

	items[5] = &item{"Golden Amulet", 250, 0, 15, 10, 10, 1, 8}
	items[6] = &item{"Iron Blade", 300, 1, 0, 60, 0, 0, 15}
	items[7] = &item{"Meat Shield", 300, 2, 100, 0, 20, 0, 0}
	items[8] = &item{"Boots of Evasion", 800, 3, 0, 0, 5, 15, 0}
	items[9] = &item{"Iron Mask", 250, 4, 40, 0, 10, 0, 0}

	items[10] = &item{"Banshee", 1000, 0, 30, 20, 10, 5, 10}
	items[11] = &item{"The unforgiven's Sword", 1000, 1, 0, 50, 5, 0, 50}
	items[12] = &item{"Dragon Skin", 1000, 2, 220, 10, 30, 0, 0}
	items[13] = &item{"Angel's Wings", 1600, 3, 30, 0, 5, 20, 0}
	items[14] = &item{"Osakar Insurance", 900, 4, 80, 0, 20, 0, 0}

	items[15] = &item{"Diamond Necklace", 2500, 0, 70, 40, 30, 8, 12}
	items[16] = &item{"Flame Axe", 3000, 1, 60, 150, 5, 0, 20}
	items[17] = &item{"Wind Wall", -1, 2, 170, 0, 100, 10, 8}
	items[18] = &item{"Phantom Dancer", 4000, 3, 0, 20, 5, 25, 1}
	items[19] = &item{"Death Cap", 2100, 4, 20, 60, 40, 0, 0}

	items[20] = &item{"Dragon slayer's Necklace", 5000, 0, 150, 80, 60, 10, 15}
	items[21] = &item{"The Sword of Lich King", -1, 1, 100, 300, 20, 0, 25}
	items[22] = &item{"Guardian Angel", 6000, 2, 500, 50, 75, 0, 0}
	items[23] = &item{"Ghost Rider", 8000, 3, 30, 20, 10, 30, 5}
	items[24] = &item{"The Helmet of Lich King", -1, 4, 320, 30, 50, 1, 5}

	items[25] = &item{"Void Spirit", 13500, 0, 300, 150, 100, 10, 20}
	items[26] = &item{"The Hardest P1nto", 15000, 1, 200, 600, 20, 0, 40}
	items[27] = &item{"Knight's Vow", 15000, 2, 1000, 100, 200, 0, 10}
	items[28] = &item{"Black Hole", -1, 3, 50, 50, 50, 35, 1}
	items[29] = &item{"Challanger's Pride", 12500, 4, 500, 50, 50, 0, 9}
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