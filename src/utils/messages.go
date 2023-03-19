package utils

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

const (
	PingEveryone = "@everyone"
	PingHere     = "@here"
)

var PingRole = regexp.MustCompile(`<&[0-9]{18}>`)

func DisablePing(msg string, users []*discordgo.User) string {
	final := msg
	for _, u := range users {
		final = strings.ReplaceAll(final, "<@"+u.ID+">", u.Username)
	}
	// Remove @everyone and @here
	final = SpeedRemover(final, PingEveryone)
	final = SpeedRemover(final, PingHere)
	// Remove @role
	final = SpeedRemoverRegex(final, PingRole)
	return final
}
