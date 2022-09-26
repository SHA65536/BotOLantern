package botolantern

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

const GLD_PATH = "guilds.json"

type BotOLantern struct {
	Session *discordgo.Session
	Guilds  *GuildStruct
	Cmds    []*discordgo.ApplicationCommand
}

type GuildStruct struct {
	Guilds map[string]bool `json:"guilds"`
	Chans  map[string]bool `json:"channels"`
	Users  map[string]bool `json:"users"`
}

func MakeHandler(token string) (*BotOLantern, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	handler := &BotOLantern{
		Session: dg,
		Guilds:  &GuildStruct{map[string]bool{}, map[string]bool{}, map[string]bool{}},
	}
	handler.LoadJson()
	dg.AddHandler(handler.slashHandler)
	dg.AddHandler(handler.messageHandler)
	return handler, nil
}

// Starts the bot
func (d *BotOLantern) Start() error {
	err := d.Session.Open()
	if err != nil {
		return fmt.Errorf("opening session: %s", err)
	}
	d.Cmds, err = d.Session.ApplicationCommandBulkOverwrite(d.Session.State.User.ID, "", Commands)
	if err != nil {
		return fmt.Errorf("creating commands: %s", err)
	}
	log.Println("[DISCORD] Bot started successfuly!")
	return nil
}

// Stops the bot
func (d *BotOLantern) Stop() error {
	log.Println("[DISCORD] Bot closing!")
	return d.Session.Close()
}

// Handles slash commands
func (d *BotOLantern) slashHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
		if i.Type == discordgo.InteractionApplicationCommand {
			log.Printf("[DISCORD] Command %s invoked", i.ApplicationCommandData().Name)
		}
		handler(d, s, i)
	}
}

// Handles chat messages
func (d *BotOLantern) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !d.Guilds.Guilds[m.GuildID] && !d.Guilds.Chans[m.ChannelID] && !d.Guilds.Users[m.Author.ID] {
		err := s.MessageReactionAdd(m.ChannelID, m.Message.ID, "🎃")
		if err != nil {
			log.Printf("Error reacting %s", err)
			return
		}
	}
}

// Load guilds json into bot
func (d *BotOLantern) LoadJson() {
	var guilds *GuildStruct = &GuildStruct{}
	jsonData, err := ioutil.ReadFile(GLD_PATH)
	if err == nil {
		err = json.Unmarshal(jsonData, guilds)
		if err == nil {
			d.Guilds = guilds
		}
	}
}

// Load guilds json into bot
func (d *BotOLantern) UpdateJson() error {
	json, err := json.MarshalIndent(d.Guilds, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(GLD_PATH, json, 0644)
	return err
}

// Pongs!
func PingCmd(d *BotOLantern, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
