package models

import "github.com/bwmarrin/discordgo"

type CommandType func(*discordgo.Session, *discordgo.MessageCreate, []string)

type Geolocate struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
		Admin1    string  `json:"admin1"`
	} `json:"results"`
}

type Weather struct {
	Current struct {
		Temperature_2m float32 `json:"temperature_2m"`
		Wind_speed_10m float32 `json:"wind_speed_10m"`
	} `json:"current"`
}

type Command struct {
	Name        string
	Description string
	Handler     CommandType
}
