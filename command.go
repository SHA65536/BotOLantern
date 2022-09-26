package botolantern

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Ping!",
	},
	{
		Name:        "lanterntoggle",
		Description: "Toggles Jack O' Lantern Reactions",
	},
}

var CommandHandlers = map[string]func(h *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":          PingCmd,
	"lanterntoggle": ToggleLantern,
}
