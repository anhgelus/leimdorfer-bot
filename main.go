package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"leimdorfer-bot/src/event"
	"leimdorfer-bot/src/utils"
	"os"
	"os/signal"
	"syscall"
)

var Discord *discordgo.Session

func main() {
	Discord, err := discordgo.New("Bot " + os.Args[1])

	Discord.AddHandler(event.HandleMessages)
	Discord.Identify.Intents = discordgo.IntentsGuildMessages

	utils.HandlePanic(err)
	err = Discord.Open()
	utils.HandlePanic(err)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	Discord.Close()
}
