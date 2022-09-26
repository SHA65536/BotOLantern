package botolantern

import "github.com/bwmarrin/discordgo"

func ToggleLantern(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Member == nil || i.Interaction.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "No permission! >:C",
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
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! I will now spook your messages!",
						Description: "🎃",
					},
				},
			},
		})
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Aye aye captain! No more spook!",
						Description: "🎃",
					},
				},
			},
		})
		return
	}
}
