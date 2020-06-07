package p1nto

import (
	"fmt"
	"strings"
	"math/rand"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	mess = make(chan *discordgo.MessageCreate, 1)
	prefix = "p."
	terminate = make(chan bool, 1)
	items = make(map[int]*item)
	players = make(map[string]*player)
	NPCs = make(map[int]*NPC)
	dungeons = make(map[int]*dungeon)
)

func Min(x, y int) int {
	if x < y {
	return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
	return x
	}
	return y
}

func RNG(x int) bool {
	//x% is the miss/crit chance
	if rand.Intn(100) < x {
		return true
	} else {
		return false
	}
}

func xRNG() int {
	return rand.Intn(100)
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	mess<- m
}

func MessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Content

	if !strings.HasPrefix(content, prefix) {
		return
	}

	content = content[len(prefix):]

	switch(strings.Split(content, " ")[0]) {
		case "die":
			if m.Author.ID == os.Getenv("OWNERID") {
				terminate<- true
				break
			}
		case "combat":
			CombatHandle(s, m)
			break
		case "farm":
			FarmHandle(s, m)
			break
		case "help":
			HelpHandle(s, m)
			break
		case "stats":
			StatsHandle(s, m)
			break
		case "equipment":
			EquipmentHandle(s, m)
			break
		case "inventory":
			InventoryHandle(s, m)
			break
		case "equip":
			EquipHandle(s, m)
			break
		case "unequip":
			UnequipHandle(s, m)
			break
		case "shop":
			ShopHandle(s, m)
			break
		case "buy":
			BuyHandle(s, m)
			break
		case "sell":
			SellHandle(s, m)
			break
		default:
			break
	}
	fmt.Println(m.Author.ID + ": " + content)
}

func Loop(s *discordgo.Session, stopListening func()) {
	InitItem()
	InitNPC()
	InitDungeon()
	LoadData()
	// defer SaveData()
	for {
		select {
			case m := <-mess:
				MessageHandle(s, m)
			case <-terminate:
				stopListening()
				fmt.Println("Terminated!")
				return
		}
	}
}