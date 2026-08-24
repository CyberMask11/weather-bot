package main

import (
	"fmt"
	"gobot/handler"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("TOKEN")

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Unable to create session.")
		return
	}

	session.AddHandler(handler.WeatherHandler)

	if err = session.Open(); err != nil {
		log.Println("Unable to open session")
	}
	defer session.Close()

	fmt.Println("Fire and Blood")

	select {}
}
