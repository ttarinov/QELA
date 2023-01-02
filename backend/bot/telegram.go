package bot

import (
	"log"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot() *TelegramBot {
	bot, err := tgbotapi.NewBotAPI("<telegram-bot-token>")
	if err != nil {
		log.Fatal(err)
	}

	return &TelegramBot{bot: bot}
}

func (b *TelegramBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I am a Telegram bot!")
			b.bot.Send(msg)
		}
	}
}
