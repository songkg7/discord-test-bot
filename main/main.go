package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load("../.env")
	println("Starting bot...")
	println(os.Getenv("DISCORD_TOKEN"))
	println("Creating session...")

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Message is created!")

	// ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("error sending message,", err)
			return
		}
	}

	if m.Content == "pong" {
		_, err := s.ChannelMessageSend(m.ChannelID, "ping")
		if err != nil {
			fmt.Println("error sending message,", err)
			return
		}
	}
}
