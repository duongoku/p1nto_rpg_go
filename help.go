package p1nto

import (
	//"fmt"
	//"strings"

	"github.com/bwmarrin/discordgo"
)

func HelpHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	temp := "```Declare prefix 'p.' to send message into bot.\nAfter prefix stating a command to play the game\n"
	temp += "- 'combat + @playername' to duel with @playername\n"
	temp += "- 'stats' to display your stats\n"
	temp += "- 'equipment' to display your equipment and Equipment ID\n"
	temp += "- 'equip' + {ID Inventory} to equip item(ID Inventory) for combat\n"
	temp += "- 'unequip' + {ID Inventory} to unequip item(ID Inventory) and set it back to your inventory\n"
	temp += "- 'shop' to go to shop and see shelves and Shop ID\n"
	temp += "- 'buy' + {Shop ID} to buy item(Shop ID) and add it to your inventory\n"
	temp += "- 'sell' + {ID Inventory} to sell item(ID Inventory) for money\n"
	temp += "- 'inventory' to show your inventory and Item ID Inventory\n"
	temp += "- 'help' to display this shit```\n"
	s.ChannelMessageSend(m.ChannelID, temp)
}
