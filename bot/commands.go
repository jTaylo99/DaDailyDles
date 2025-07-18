package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "dlelist",
		Description: "Get the full list of dles",
	},
}

func registerCommands(session *discordgo.Session, guildID string) {
	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
