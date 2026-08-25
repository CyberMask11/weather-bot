package handler

import (
	"gobot/commands"
	"gobot/models"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		for name, info := range Commands {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title: name,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  info.Name,
						Value: info.Description,
					},
				},
				Color: 0xFF0000,

				Footer: &discordgo.MessageEmbedFooter{
					Text: m.Author.DisplayName(),
				},
				Timestamp: time.Now().Format(time.RFC3339),
			})
		}
		return
	}

	return
}

var Commands = map[string]models.Command{
	"!weather": {
		Name:        "Weather Command.",
		Description: "Get the weather of a city. weather cityname",
		Handler:     commands.Weather,
	},
	"!hello": {
		Name:        "Hello Command",
		Description: "The Bot greets you.",
		Handler:     commands.Hello,
	},
}

func WeatherHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.ChannelID != "1539982012415348776" {
		return
	}

	fields := strings.Fields(m.Content)

	if m.Content == "!help" {
		helpCommand(s, m, fields[1:])
		return
	}

	commandFunc, ok := Commands[fields[0]]
	if !ok {
		s.ChannelMessageSend(m.ChannelID, "Invalid command, !help for info")
		return
	}

	commandFunc.Handler(s, m, fields[1:])
}
