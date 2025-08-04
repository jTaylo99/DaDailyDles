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
	/*
		This function is for handling "/" commands. We are able to define new / commands by adding onto this switch statement. NOTE: Commands need to also be registered in commands.go as this is what sends them to the Discord API.
	*/
	switch i.ApplicationCommandData().Name {
	case "dlelist":
		dleList(s, i)
	}
}

func Run(BotToken string, GuildIDs []string) {
	discord, err := discordgo.New("Bot " + BotToken) // Given Go is a typed language, we use := to declare our vars, and the function they call decides their type.
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
