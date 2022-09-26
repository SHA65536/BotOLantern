package botolantern

import (
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
	d.Guilds.Guilds[i.Interaction.GuildID] = !d.Guilds.Guilds[i.Interaction.GuildID]
	d.UpdateJson()
	if !d.Guilds.Guilds[i.Interaction.GuildID] {
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
