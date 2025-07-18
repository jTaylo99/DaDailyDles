package bot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type dle struct {
	name        string
	description string
	updateTime  *time.Time
	url         string
}

var loc *time.Location
var dles []dle

func init() {
	var err error
	loc, err = time.LoadLocation("Pacific/Auckland")
	if err != nil {
		log.Fatalf("Failed to load timezone: %v", err)
	}

	dles = []dle{
		{
			name:        "Scrandle",
			description: "Guess which stadium food is most popular",
			updateTime:  nil,
			url:         "https://scrandle.com",
		},
		{
			name:        "Bandle",
			description: "Guess the song, one instrument at a time",
			updateTime:  nil,
			url:         "https://bandle.app/daily",
		},
		{
			name:        "GuessTheGame",
			description: "Guess a game based off screenshots",
			updateTime: func() *time.Time {
				now := time.Now().In(loc)
				t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
				return &t
			}(),
			url: "https://guessthe.game/",
		},
	}
}

func generateDleListMessage() string {
	var sb strings.Builder
	now := time.Now().In(loc)

	for _, dle := range dles {
		nextUpdate := "-"
		if dle.updateTime != nil {
			// Ensure update time is set to today's date
			t := time.Date(now.Year(), now.Month(), now.Day(),
				dle.updateTime.Hour(), dle.updateTime.Minute(), 0, 0, loc)
			if now.After(t) {
				t = t.Add(24 * time.Hour)
			}
			duration := t.Sub(now).Round(time.Second)
			nextUpdate = duration.String()
		} else {
			nextUpdate = "Device dependent"
		}

		sb.WriteString(fmt.Sprintf(
			"**%s** – %s\n🕐 Next update: `%s`\n[▶️ Play](%s)\n\n",
			dle.name, dle.description, nextUpdate, dle.url,
		))
	}

	return sb.String()
}

func dleList(s *discordgo.Session, i *discordgo.InteractionCreate) {
	table := generateDleListMessage()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: table,
			Flags:   discordgo.MessageFlagsSuppressEmbeds,
		},
	})
	if err != nil {
		log.Println("Failed to send interaction response:", err)
	}
}
