package event

import (
	"github.com/bwmarrin/discordgo"
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
	rand.NewSource(time.Now().UnixNano())
	v := rand.Intn(len(Replacers))
	msg := utils.NewCaseInsensitiveReplacer("du coup", Replacers[v]).Replace(m.Content)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	utils.HandlePanic(err)
}
