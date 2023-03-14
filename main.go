package main

import (
	"github.com/bwmarrin/discordgo"
	"leimdorfer-bot/src/utils"
	"os"
)

var Discord *discordgo.Session

func main() {
	Discord, err := discordgo.New("Bot " + os.Args[1])
	utils.HandlePanic(err)
	err = Discord.Open()
	utils.HandlePanic(err)
}
