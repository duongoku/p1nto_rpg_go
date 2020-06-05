package p1nto

import (
	//"fmt"
	//"strings"

	"github.com/bwmarrin/discordgo"
)

func HelpHandle(s *discordgo.Session, m *discordgo.MessageCreate){
	temp := "```Declare prefix 'p.' to send message into bot.\nAfter prefix stating a command to play the game\n"
	temp += "- 'combat + @playername' to duel with @playername\n"
	temp += "- 'stats' to display your stats\n"
	temp += "- 'equipment' to display your equipment and Equipment ID\n"
	temp += "- 'equip' + 'Equipment ID' to equip item(Equipment ID) for combat\n"
	temp += "- 'unequip' + 'Equipment ID' to unequip item(Equipment ID) and set it back to your inventory\n"
	temp += "- 'shop' to go to shop and see shelves and Item ID\n"
	temp += "- 'buy' + 'Item ID' to buy item(Item ID) and add it to your inventory\n"
	temp += "- 'sell' + 'Item ID' to sell item(Item ID) for money\n"
	temp += "- 'inventory' to show your inventory\n"
	temp += "- 'help' to display this shit```\n"
	s.ChannelMessageSend(m.ChannelID, temp)
}