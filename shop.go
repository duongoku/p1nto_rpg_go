package p1nto

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"fmt"
	"strings"
)

func ShopHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	temp := "```Welcome to the shop\n"
	for itID, eqm := range items {
		temp = temp + "\n" + "Shop ID:" + strconv.Itoa(itID) + " Price: " + strconv.Itoa(eqm.Price) + "$\n"
		temp += eqm.Name + "-Type:" + strconv.Itoa(eqm.SlotID) + "-"
		if eqm.Hp > 0 {
			temp = temp + "Hp+" + strconv.Itoa(eqm.Hp) + ", "
		}
		if eqm.Atk > 0 {
			temp = temp + "Attack+" + strconv.Itoa(eqm.Atk) + ", "
		}
		if eqm.Def > 0 {
			temp = temp + "Defense+" + strconv.Itoa(eqm.Def) + ", "
		}
		if eqm.Evasion > 0 {
			temp = temp + "Evasion+" + strconv.Itoa(eqm.Evasion) + "%, "
		}
		if eqm.CritChance > 0 {
			temp = temp + "Critical Chance+" + strconv.Itoa(eqm.CritChance) + "%, "
		}
		temp = temp[0:len(temp)-2]
		temp+="\n"
	}
	temp += "```"
	s.ChannelMessageSend(m.ChannelID, temp)
}

func BuyHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	subcontent := strings.Split(m.Content, " ")
	if len(subcontent) < 1 {
		s.ChannelMessageSend(m.ChannelID, "You must mention an itemID in the shop !")
		return
	}

	itemID, _ := strconv.Atoi(subcontent[len(subcontent) - 1])

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

func SellHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	subcontent := strings.Split(m.Content, " ")
	invID, _ := strconv.Atoi(subcontent[len(subcontent) - 1])
	fmt.Println(invID)

	u := m.Author
	CheckPlayer(u)

	if invID >= len(users[u.ID].Inventory) || invID < 0 {
		s.ChannelMessageSend(m.ChannelID, "There is no such item!")
	} else {
		users[u.ID].Money += items[users[u.ID].Inventory[invID]].Price

		temp := items[users[u.ID].Inventory[invID]].Name + " sold!\n"
		temp += "You get " + strconv.Itoa(items[users[u.ID].Inventory[invID]].Price) + "$ back"

		s.ChannelMessageSend(m.ChannelID, temp)
		users[u.ID].Inventory = append(users[u.ID].Inventory[:invID], users[u.ID].Inventory[invID+1:]...)
	}

}