package utils

import "github.com/bwmarrin/discordgo"

func GetUser(dg *discordgo.Session, guildId string, userId string) *discordgo.Member {
	g, err := dg.Guild(guildId)
	HandlePanic(err)
	for _, member := range g.Members {
		if member.User.ID == userId {
			return member
		}
	}
	return nil
}
