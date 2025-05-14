package discord

import (
	"fmt"
	"log"

	"github.com/ADG08/tracker-lol/domain/ports"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(token string) (ports.DiscordService, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages

	bot := &Bot{session: session}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Bot connecté en tant que %s", r.User.Username)
	})

	if err := bot.session.Open(); err != nil {
		return nil, err
	}

	return bot, nil
}

func CloseBot(bot ports.DiscordService) error {
	if b, ok := bot.(*Bot); ok {
		return b.session.Close()
	}
	return nil
}

func (b *Bot) SendMessage(channelID, message string) error {
	_, err := b.session.ChannelMessageSend(channelID, message)
	return err
}

func (b *Bot) RegisterCommand(name, description string, options []ports.CommandOption) error {
	discordOptions := make([]*discordgo.ApplicationCommandOption, len(options))
	for i, opt := range options {
		discordOptions[i] = &discordgo.ApplicationCommandOption{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        opt.Name,
			Description: opt.Description,
			Required:    opt.Required,
		}
	}

	cmd := &discordgo.ApplicationCommand{
		Name:        name,
		Description: description,
		Options:     discordOptions,
	}

	appID := b.session.State.User.ID
	guilds := b.session.State.Guilds
	if len(guilds) == 0 {
		return fmt.Errorf("bot is not connected to any server")
	}

	guildID := guilds[0].ID

	commands, err := b.session.ApplicationCommands(appID, guildID)
	if err != nil {
		return fmt.Errorf("failed to fetch existing commands: %v", err)
	} else {
		for _, existingCmd := range commands {
			if existingCmd.Name == name {
				err := b.session.ApplicationCommandDelete(appID, guildID, existingCmd.ID)
				if err != nil {
					return fmt.Errorf("failed to delete existing command: %v", err)
				}
			}
		}
	}

	_, err = b.session.ApplicationCommandCreate(appID, guildID, cmd)
	if err != nil {
		return fmt.Errorf("failed to create command: %v", err)
	}

	return nil
}

func (b *Bot) HandleCommand(name string, handler func(interaction *ports.Interaction) error) error {
	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if i.ApplicationCommandData().Name != name {
			return
		}

		options := make(map[string]string)
		for _, opt := range i.ApplicationCommandData().Options {
			options[opt.Name] = opt.StringValue()
		}

		interaction := &ports.Interaction{
			UserID:    i.Member.User.ID,
			ChannelID: i.ChannelID,
			Options:   options,
		}

		if err := handler(interaction); err != nil {
			log.Printf("Erreur lors du traitement de la commande: %v", err)
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: err.Error(),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				log.Printf("Erreur lors de l'envoi du message d'erreur: %v", err)
			}
			return
		}
	})
	return nil
}
