package services

import (
	"context"
	"fmt"

	"github.com/ADG08/tracker-lol/domain/model"
	"github.com/ADG08/tracker-lol/domain/ports"
	"github.com/ADG08/tracker-lol/domain/repository"
)

type CommandService struct {
	discordService ports.DiscordService
	userRepository repository.UserRepository
}

func NewCommandService(discordService ports.DiscordService, userRepository repository.UserRepository) *CommandService {
	return &CommandService{
		discordService: discordService,
		userRepository: userRepository,
	}
}

func (s *CommandService) RegisterCommands() error {
	if err := s.discordService.RegisterCommand(
		"register",
		"Enregistrer votre compte Riot",
		[]ports.CommandOption{
			{
				Name:        "pseudo",
				Description: "Votre pseudo Riot",
				Required:    true,
			},
			{
				Name:        "tag",
				Description: "Votre tag Riot (ex: EUW)",
				Required:    true,
			},
		},
	); err != nil {
		return err
	}

	if err := s.discordService.RegisterCommand(
		"profil",
		"Afficher votre profil Riot",
		[]ports.CommandOption{},
	); err != nil {
		return err
	}

	return nil
}

func (s *CommandService) HandleRegisterCommand(interaction *ports.Interaction) error {
	existingUser, err := s.userRepository.FindFromDiscordID(interaction.UserID, context.Background())
	if err == nil && existingUser != nil {
		return fmt.Errorf("vous êtes déjà enregistré avec le pseudo %s et le tag %s", existingUser.GetPseudo(), existingUser.GetTag())
	}

	user := model.CreateUser(
		interaction.UserID,
		interaction.Options["pseudo"],
		interaction.Options["tag"],
	)

	_, err = s.userRepository.Persist(user, context.Background())
	if err != nil {
		return fmt.Errorf("erreur lors de l'enregistrement : %v", err)
	}

	return fmt.Errorf("compte enregistré avec succès !\nPseudo : %s\nTag : %s", user.GetPseudo(), user.GetTag())
}

func (s *CommandService) HandleProfilCommand(interaction *ports.Interaction) error {
	user, err := s.userRepository.FindFromDiscordID(interaction.UserID, context.Background())
	if err != nil {
		return fmt.Errorf("vous n'êtes pas encore enregistré. Utilisez la commande /register pour vous enregistrer")
	}

	return fmt.Errorf("votre profil Riot :\nPseudo : %s\nTag : %s", user.GetPseudo(), user.GetTag())
}
