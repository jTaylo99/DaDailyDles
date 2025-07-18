package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatalf("Error creating Discord session: %v", e)
	}
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "dlelist":
		dleList(s, i)
	}
}

func Run(BotToken string, GuildIDs []string) {
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	discord.AddHandler(interactionHandler)

	discord.Open()
	defer discord.Close()

	for _, guildID := range GuildIDs {
		registerCommands(discord, guildID)
	}

	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
