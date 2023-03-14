package users

import "github.com/bwmarrin/discordgo"

type User struct {
	*discordgo.User
	GuildID string
	Warning int
}
