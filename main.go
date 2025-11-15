package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error Creating Discord Session!", err)
		return
	} else {
		fmt.Println("Discord Session Found! Tight")
	}

	discord.AddHandler(roastmeCreate)

	discord.AddHandler(roastyouCreate)

	discord.AddHandler(mentionsRoastYouCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error Opening Connection", err)
	}

	fmt.Println("Bot is now running, press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()

}

