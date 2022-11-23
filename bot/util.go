package bot

import (
	"dustin-ward/AdventOfCodeBot/data"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func getChannel(guildId string) (*data.Channel, error) {
	for _, ch := range c {
		if ch.GuildId == guildId {
			return &ch, nil
		}
	}
	return nil, fmt.Errorf("Error: channel not found")
}

func respondWithError(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})

	if err != nil {
		log.Fatal(fmt.Errorf("Fatal: %v", err))
	}
}
