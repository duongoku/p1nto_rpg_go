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
	temp += "- 'equip' + {Inventory ID} to equip item(Inventory ID) for combat\n"
	temp += "- 'unequip' + {Inventory ID} to unequip item(Inventory ID) and set it back to your inventory\n"
	temp += "- 'shop' to go to shop and see shelves and Shop ID\n"
	temp += "- 'buy' + {Shop ID} to buy item(Shop ID) and add it to your inventory\n"
	temp += "- 'sell' + {Inventory ID} to sell item(Inventory ID) for money\n"
	temp += "- 'inventory' to show your inventory and Item Inventory ID\n"
	temp += "- 'help' to display this shit```\n"
	s.ChannelMessageSend(m.ChannelID, temp)
}
