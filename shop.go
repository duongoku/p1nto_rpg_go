package p1nto

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"fmt"
	"strings"
)

func ShopHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	temp := "```Welcome to the shop"
	total := len(items)
	for i := 0; i < total; i++ {
		temp = temp + "\n\nItem ID:" + strconv.Itoa(i) + " Price: $" + strconv.Itoa(items[i].Price)
		temp = temp + "\n" + items[i].Name + " | Equip Slot:" + strconv.Itoa(items[i].SlotID) + " | "
		if items[i].Hp > 0 {
			temp = temp + "Hp+" + strconv.Itoa(items[i].Hp) + ", "
		}
		if items[i].Atk > 0 {
			temp = temp + "Attack+" + strconv.Itoa(items[i].Atk) + ", "
		}
		if items[i].Def > 0 {
			temp = temp + "Defense+" + strconv.Itoa(items[i].Def) + ", "
		}
		if items[i].Evasion > 0 {
			temp = temp + "Evasion+" + strconv.Itoa(items[i].Evasion) + "%, "
		}
		if items[i].CritChance > 0 {
			temp = temp + "Critical Chance+" + strconv.Itoa(items[i].CritChance) + "%, "
		}
		temp = temp[0:len(temp)-2]
	}
	temp += "```"
	s.ChannelMessageSend(m.ChannelID, temp)
}

func BuyHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide an item ID showed in the shop!")
		return
	}
	if len(content) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}


	itemID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong Format")
  		fmt.Println(err)
  		return
	}

	u := m.Author
	CheckPlayer(u)

	if itemID >= len(items) || itemID < 0 {
		s.ChannelMessageSend(m.ChannelID, "There is no such item!")
	} else if items[itemID].Price > users[u.ID].Money {
		s.ChannelMessageSend(m.ChannelID, "Not enough money you poor litte shit :smirk: !")
	} else {
		users[u.ID].Inventory = append(users[u.ID].Inventory, itemID)
		users[u.ID].Money -= items[itemID].Price
		s.ChannelMessageSend(m.ChannelID, items[itemID].Name + " has been added to your Inventory !")
	}
}

func SellHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Split(m.Content, " ")
	if len(content) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide an ID in Inventory in the shop!")
		return
	}
	if len(content) > 2 {
		s.ChannelMessageSend(m.ChannelID, "Too many arguments")
		return
	}

	invID, err := strconv.Atoi(content[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Wrong Format")
  		fmt.Println(err)
  		return
	}

	u := m.Author
	CheckPlayer(u)

	if invID >= len(users[u.ID].Inventory) || invID < 0 {
		s.ChannelMessageSend(m.ChannelID, "There is no such item!")
	} else {
		users[u.ID].Money += items[users[u.ID].Inventory[invID]].Price

		temp := items[users[u.ID].Inventory[invID]].Name + "is sold!\n"
		temp += "You get " + strconv.Itoa(items[users[u.ID].Inventory[invID]].Price) + "$ back"

		s.ChannelMessageSend(m.ChannelID, temp)
		users[u.ID].Inventory = append(users[u.ID].Inventory[:invID], users[u.ID].Inventory[invID+1:]...)
	}

}