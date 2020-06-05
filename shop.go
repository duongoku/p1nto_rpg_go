package p1nto

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"fmt"
	"strings"
)

func ShopHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	temp := "```Value type in order: Hp, Price, Atk, Def, Evasion, CritChance\n"
	for i, value := range items {
		temp += "Slot " + strconv.Itoa(i) + ": "
		temp += value.Name + ": " + "\n"
		temp += strconv.Itoa(value.Price) + " "
		temp += strconv.Itoa(value.Hp) + " "
		temp += strconv.Itoa(value.Atk) + " "
		temp += strconv.Itoa(value.Def) + " "
		temp += strconv.Itoa(value.Evasion) + " "
		temp += strconv.Itoa(value.CritChance) + "\n"
	}
	temp += "```"
	s.ChannelMessageSend(m.ChannelID, temp)
}

func BuyHandle(content string, s *discordgo.Session, m *discordgo.MessageCreate){
	subcontent := strings.Split(content, " ")
	itemID, _ := strconv.Atoi(subcontent[len(subcontent) - 1])
	fmt.Println(itemID)

	u := m.Author
	CheckPlayer(u)

	if itemID > len(items) {
		s.ChannelMessageSend(m.ChannelID, "There is no such item!")
	} else if items[itemID].Price > users[u.ID].Money {
		s.ChannelMessageSend(m.ChannelID, "Not enough money you poor litte shit :smirk: !")
	} else {
		users[u.ID].Inventory = append(users[u.ID].Inventory, items[itemID])
		users[u.ID].Money -= items[itemID].Price
		s.ChannelMessageSend(m.ChannelID, items[itemID].Name + " has been added to your Inventory !")
	}
}

func SellHandle(content string, s *discordgo.Session, m *discordgo.MessageCreate){
	subcontent := strings.Split(content, " ")
	invID, _ := strconv.Atoi(subcontent[len(subcontent) - 1])
	fmt.Println(invID)

	u := m.Author
	CheckPlayer(u)

	if invID > len(users[u.ID].Inventory) {
		s.ChannelMessageSend(m.ChannelID, "There is no such item!")
	} else {
		users[u.ID].Money += users[u.ID].Inventory[invID].Price

		tmps := users[u.ID].Inventory[invID].Name + "sold!\n"
		tmps += "You get " + strconv.Itoa(users[u.ID].Inventory[invID].Price) + "$ back"

		s.ChannelMessageSend(m.ChannelID, tmps)
		users[u.ID].Inventory = append(users[u.ID].Inventory[:invID], users[u.ID].Inventory[invID + 1:]...)
	}

}