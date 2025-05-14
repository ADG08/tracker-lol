package discord

import (
	"github.com/ADG08/tracker-lol/application/services"
	"github.com/ADG08/tracker-lol/domain/ports"
	"github.com/ADG08/tracker-lol/domain/repository"
)

type Handler struct {
	commandService *services.CommandService
	discordService ports.DiscordService
}

func NewHandler(discordService ports.DiscordService, userRepository repository.UserRepository) *Handler {
	return &Handler{
		commandService: services.NewCommandService(discordService, userRepository),
		discordService: discordService,
	}
}

func (h *Handler) Start() error {
	if err := h.commandService.RegisterCommands(); err != nil {
		return err
	}

	if err := h.discordService.HandleCommand("register", h.commandService.HandleRegisterCommand); err != nil {
		return err
	}

	if err := h.discordService.HandleCommand("profil", h.commandService.HandleProfilCommand); err != nil {
		return err
	}

	return nil
}
