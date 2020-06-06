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
	temp += "- 'equip' + {ID in Inventory} to equip item(ID in Inventory) for combat\n"
	temp += "- 'unequip' + {ID in Inventory} to unequip item(ID in Inventory) and set it back to your inventory\n"
	temp += "- 'shop' to go to shop and see shelves and Shop ID\n"
	temp += "- 'buy' + {Shop ID} to buy item(Shop ID) and add it to your inventory\n"
	temp += "- 'sell' + {ID in Inventory} to sell item(ID in Inventory) for money\n"
	temp += "- 'inventory' to show your inventory and Item ID in Inventory\n"
	temp += "- 'help' to display this shit```\n"
	s.ChannelMessageSend(m.ChannelID, temp)
}
