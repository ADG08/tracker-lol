package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gihtub.com/ADG08/tracker-lol/app/internal/config"
	"gihtub.com/ADG08/tracker-lol/app/internal/discord"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	bot, err := discord.NewBot(cfg.DiscordToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	defer bot.Close()

	if err := bot.SendMessage(cfg.ChannelID, "Hello, world!"); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}
