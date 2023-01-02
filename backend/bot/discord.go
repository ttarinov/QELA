package bot

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	Session *discordgo.Session
}

func NewDiscordBot() *DiscordBot {
	session, err := discordgo.New("<your-discord-bot-token>")
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		os.Exit(1)
	}

	return &DiscordBot{Session: session}
}

func (b *DiscordBot) Start() {
	b.Session.AddHandler(onMessage)

	err := b.Session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		os.Exit(1)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I am a Discord bot! Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	b.Session.Close()
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, I am a Discord bot!")
	}
}
