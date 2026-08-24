package commands

import (
	"encoding/json"
	"fmt"
	"gobot/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Geolocate(city string) ([]float64, []string) {
	var geolocate models.Geolocate

	resp, err := http.Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s", city))
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	if err = json.Unmarshal(body, &geolocate); err != nil {
		log.Println(err)
		return nil, nil
	}

	data := []float64{
		geolocate.Results[0].Latitude,
		geolocate.Results[0].Longitude,
	}

	citydata := []string{
		city,
		geolocate.Results[0].Admin1,
		geolocate.Results[0].Country,
	}

	return data, citydata
}

func Weather(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var weather models.Weather

	data, city := Geolocate(args[0])

	resp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m,wind_speed_10m", data[0], data[1]))
	if err != nil {
		log.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if err = json.Unmarshal(body, &weather); err != nil {
		log.Println(err)
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("📍 %s", city[0]),
		Description: fmt.Sprintf("🌤️ %s, %s", city[1], city[2]),
		Fields: []*discordgo.MessageEmbedField{
			{
				Value: fmt.Sprintf("🌡️ %v °C", weather.Current.Temperature_2m),
			},
			{
				Value: fmt.Sprintf("💨 %v km/h", weather.Current.Wind_speed_10m),
			},
		},

		Color: 0x800080,

		Footer: &discordgo.MessageEmbedFooter{
			Text: m.Author.DisplayName(),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
