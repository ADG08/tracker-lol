package repository

import (
	"context"

	"github.com/ADG08/tracker-lol/domain/model"
)

type UserRepository interface {
	Persist(user *model.User, ctx context.Context) (*model.User, error)
	FindFromDiscordID(discordID string, ctx context.Context) (*model.User, error)
}
