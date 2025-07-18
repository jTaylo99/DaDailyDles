package main

import (
	"os"

	"github.com/jTaylo99/DaDailyDles/bot"
)

func main() {
	GuildIDs := []string{"1186255506671673384"}
	bot.Run(os.Getenv("discordToken"), GuildIDs)
}
