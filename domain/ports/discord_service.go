// domain/ports/discord_service.go
package ports

type DiscordService interface {
	SendMessage(channelID, message string) error
	HandleCommand(name string, handler func(interaction *Interaction) error) error
	RegisterCommand(name, description string, options []CommandOption) error
}

type CommandOption struct {
	Name        string
	Description string
	Required    bool
}

type Interaction struct {
	UserID    string
	ChannelID string
	Options   map[string]string
}
