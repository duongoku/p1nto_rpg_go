package p1nto

import (
	//"fmt"
	//"strings"

	"github.com/bwmarrin/discordgo"
)

func HelpHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	temp := "```"
	temp += "Declare prefix 'p.' to send message into bot.\nAfter prefix stating a command to play the game\n"
	temp += "- 'help' to display this\n"
	temp += "- 'combat + @playername' to duel with @playername\n"
	temp += "- 'stats' to display your stats\n"
	temp += "- 'inventory' to show your inventory\n"
	temp += "- 'equipment' to display your equipments details\n"
	temp += "- 'equip' + {ID in Inventory} to equip an item\n"
	temp += "- 'unequip' + {Slot ID} to unequip an item\n"
	temp += "- 'shop' to go to shop and see shelves\n"
	temp += "- 'buy' + {Item ID} to buy an item\n"
	temp += "- 'sell' + {ID in Inventory} to sell an item\n"
	temp += "- 'dungeon' + {Dungeon ID} to see a dungeon's monsters\n"
	temp += "- 'npc' + {NPC ID} to see a monster/NPC details\n"
	temp += "- 'farm' + {Dungeon ID} to farm a dungeon\n"
	temp += "```"
	s.ChannelMessageSend(m.ChannelID, temp)
}
