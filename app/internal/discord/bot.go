package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		session: session,
	}

	if err := bot.session.Open(); err != nil {
		return nil, err
	}

	log.Println("Bot is running")
	return bot, nil
}

func (b *Bot) SendMessage(channelID, message string) error {
	_, err := b.session.ChannelMessageSend(channelID, message)
	return err
}

func (b *Bot) Close() {
	b.session.Close()
}
