package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ADG08/tracker-lol/config"
	"github.com/ADG08/tracker-lol/infrastructure/discord"
	"github.com/ADG08/tracker-lol/infrastructure/persistence"
	"github.com/ADG08/tracker-lol/infrastructure/persistence/repository"
	interfaceDiscord "github.com/ADG08/tracker-lol/interfaces/discord"
)

func main() {
	log.Println("Démarrage du bot...")

	discordCfg, dbCfg, err := config.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement de la configuration:", err)
	}
	log.Println("Configuration chargée avec succès")

	db := persistence.New(dbCfg)
	if !db.Health() {
		log.Fatal("La base de données n'est pas accessible")
	}
	log.Println("Connexion à la base de données établie")

	bot, err := discord.NewBot(discordCfg.DiscordToken)
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation du bot Discord:", err)
	}
	defer discord.CloseBot(bot)
	log.Println("Bot Discord initialisé")

	userRepo := repository.NewUserRepository()
	log.Println("Repositories initialisés")

	handler := interfaceDiscord.NewHandler(bot, userRepo)
	if err := handler.Start(); err != nil {
		log.Fatal("Erreur lors du démarrage du handler:", err)
	}
	log.Println("Handler démarré avec succès")

	if err := bot.SendMessage(discordCfg.ChannelID, "Bot initialisé et prêt à recevoir des commandes"); err != nil {
		log.Printf("Impossible d'envoyer le message de confirmation: %v", err)
	}

	log.Println("Bot en attente de commandes...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	log.Println("Arrêt du bot...")
}
