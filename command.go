package botolantern

import "github.com/bwmarrin/discordgo"

var minChance float64 = 0

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Ping!",
	},
	{
		Name:        "lanterntoggle",
		Description: "Toggles Jack O' Lantern Reactions In The Whole Server",
	},
	{
		Name:        "channeltoggle",
		Description: "Toggles Jack O' Lantern Reactions In This Channel",
	},
	{
		Name:        "personaltoggle",
		Description: "Toggles Jack O' Lantern Reactions For User",
	},
	{
		Name:        "lanternchance",
		Description: "Sets lantern chance",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "chance",
				Description: "Chance",
				MinValue:    &minChance,
				MaxValue:    100,
				Required:    true,
			},
		},
	},
}

var CommandHandlers = map[string]func(h *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":           PingCmd,
	"lanterntoggle":  ToggleLantern,
	"channeltoggle":  RestrictChannel,
	"personaltoggle": RestrictUser,
	"lanternchance":  LanternChance,
}
