package utils

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func DisablePing(msg string, users []*discordgo.User) string {
	final := msg
	for _, u := range users {
		final = strings.ReplaceAll(final, "<@"+u.ID+">", u.Username)
	}
	return final
}
