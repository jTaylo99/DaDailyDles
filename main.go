package main

// Import bot dependencies

/*
We are importing the repository itself as they are considered seperate packages, even if they are in the same repository; TODO: Check best practice, this feels weird?
*/
import (
	"os"

	"github.com/jTaylo99/DaDailyDles/bot"
)

func main() {
	GuildIDs := []string{"1186255506671673384"} // GuildID represents Discord Servers; TODO: Have different way of uploading API commands to server.
	bot.Run(os.Getenv("discordToken"), GuildIDs)
}
