package users

import "github.com/bwmarrin/discordgo"

var users = map[string]User{}

func GetUser(u *discordgo.User, g string) *User {
	id := u.ID + "|" + g
	user, ok := users[id]
	if !ok {
		users[id] = user
		user = User{Warning: 0, User: u, GuildID: g}
	}
	return &user
}

func (u *User) Save() User {
	users[u.ID+"|"+u.GuildID] = *u
	return *u
}
