# gobot

A Discord bot written in Go. Currently features a weather command powered by the [Open-Meteo API](https://open-meteo.com/).

## Features

- `!weather <city>` — Get current temperature and wind speed for a city
- `!help` — List available commands

## Tech Stack

- [Go](https://go.dev/)
- [discordgo](https://github.com/bwmarrin/discordgo)
- [Open-Meteo API](https://open-meteo.com/) (weather + geocoding)

## Getting Started

### Prerequisites

- Go 1.26+
- A Discord bot token (see [creating a bot application](https://discord.com/developers/applications))

### Installation

1. Clone the repository

```bash
git clone https://github.com/<your-username>/gobot.git
cd gobot
```

2. Create a `.env` file in the project root:

```env
TOKEN=your-discord-bot-token
```

3. Run the bot

```bash
go run main.go
```

## Adding a New Command

1. Implement the command function in `commands/`
2. Register it in the `Commands` map in `handler/weatherHandler.go`

## License

MIT
