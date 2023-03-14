package event

import "github.com/bwmarrin/discordgo"

func HandleMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}
