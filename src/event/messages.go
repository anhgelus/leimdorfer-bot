package event

import (
	"github.com/bwmarrin/discordgo"
	"leimdorfer-bot/src/users"
	"leimdorfer-bot/src/utils"
	"math/rand"
	"strings"
	"time"
)

var (
	Replacers = [8]string{"ainsi", "de ce fait", "en conséquence", "conséquemment", "subséquement", "ipso facto",
		"de sorte que", "corrolairement"}
)

func HandleMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "du coup") {
		handleDuCoup(s, m)
	}
}

func handleDuCoup(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := m.GuildID
	rand.NewSource(time.Now().UnixNano())
	v := rand.Intn(len(Replacers))
	msg := utils.NewCaseInsensitiveReplacer("du coup", Replacers[v]).Replace(m.Content)
	ref := discordgo.MessageReference{MessageID: m.ID, ChannelID: m.ChannelID, GuildID: guildID}
	_, err := s.ChannelMessageSendReply(m.ChannelID, msg, &ref)
	utils.HandlePanic(err)
	// Get the user and update the warning counter
	user := users.GetUser(m.Author, m.GuildID)
	user.Warning += 1
	user.Save()

	// Check if we must time out the user
	if user.Warning != 5 {
		return
	}
	// Reset the warning count
	user.Warning = 0
	user.Save()
	// Send the informative message
	_, err = s.ChannelMessageSendReply(m.ChannelID, "**Get Leimdorfered!**\nTu viens de te faire timeout pour 60 secondes !", &ref)
	utils.HandlePanic(err)
	// Calc the time
	d := 60 * time.Second
	t := time.Now().Add(d)
	// Time out the user
	err = s.GuildMemberTimeout(guildID, user.ID, &t)
	utils.HandlePanic(err)
}
