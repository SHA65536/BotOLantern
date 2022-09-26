package botolantern

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ToggleLantern(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Member == nil || i.Interaction.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "No permission! >:C",
						Color: 0xFF0000,
					},
				},
			},
		})
		return
	}
	if d.Guilds.Guilds[i.Interaction.GuildID] == 0 {
		d.Guilds.Guilds[i.Interaction.GuildID] = 100
	} else {
		d.Guilds.Guilds[i.Interaction.GuildID] = 0
	}
	d.UpdateJson()
	if d.Guilds.Guilds[i.Interaction.GuildID] == 100 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! I will now spook your messages!",
						Description: "🎃",
						Color:       0x00FF00,
					},
				},
			},
		})
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! No more spook!",
						Description: "🎃",
						Color:       0xFF0000,
					},
				},
			},
		})
		return
	}
}

func RestrictChannel(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Member == nil || i.Interaction.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "No permission! >:C",
						Color: 0xFF0000,
					},
				},
			},
		})
		return
	}
	d.Guilds.Chans[i.Interaction.ChannelID] = !d.Guilds.Chans[i.Interaction.ChannelID]
	d.UpdateJson()
	if !d.Guilds.Chans[i.Interaction.ChannelID] {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! I will now spook this channel!",
						Description: "🎃",
						Color:       0x00FF00,
					},
				},
			},
		})
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! No more spook in this channel!",
						Description: "🎃",
						Color:       0xFF0000,
					},
				},
			},
		})
		return
	}
}

func RestrictUser(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Member == nil {
		return
	}
	d.Guilds.Users[i.Interaction.Member.User.ID] = !d.Guilds.Users[i.Interaction.Member.User.ID]
	d.UpdateJson()
	if !d.Guilds.Users[i.Interaction.Member.User.ID] {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! I will now spook you!",
						Description: "🎃",
						Color:       0x00FF00,
					},
				},
			},
		})
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! No more spook for you!",
						Description: "🎃",
						Color:       0xFF0000,
					},
				},
			},
		})
		return
	}
}

func LanternChance(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Member == nil || i.Interaction.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "No permission! >:C",
						Color: 0xFF0000,
					},
				},
			},
		})
		return
	}
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	d.Guilds.Guilds[i.Interaction.GuildID] = optionMap["chance"].IntValue()
	d.UpdateJson()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("Aye aye captain! Set chance to %d", optionMap["chance"].IntValue()),
					Description: "🎃",
					Color:       0x00FF00,
				},
			},
		},
	})
}
