package main

import (
	"fmt"

	"github.com/ttarinov/qela/backend/bot"
	"github.com/ttarinov/qela/backend/nlp"
)

func main() {
	// Initialize the Telegram and Discord bots
	telegramBot := bot.NewTelegramBot()
	discordBot := bot.NewDiscordBot()

	// Start the bots
	go telegramBot.Start()
	go discordBot.Start()

	// Run some text analysis
	result := nlp.AnalyzeText("This is some text to analyze")
	fmt.Println(result)
}
