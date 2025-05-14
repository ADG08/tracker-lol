package repository

import (
	"context"

	"github.com/ADG08/tracker-lol/domain/model"
	"github.com/ADG08/tracker-lol/infrastructure/persistence"
	db "github.com/ADG08/tracker-lol/infrastructure/persistence/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		q: db.New(persistence.Conn()),
	}
}

func (r *UserRepository) Persist(user *model.User, ctx context.Context) (*model.User, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(user.GetID()); err != nil {
		return nil, err
	}
	params := db.PersistUserParams{
		ID:        uuid,
		DiscordID: user.GetDiscordID(),
		Pseudo:    user.GetPseudo(),
		Tag:       user.GetTag(),
	}
	err := r.q.PersistUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindFromDiscordID(discordID string, ctx context.Context) (*model.User, error) {
	data, err := r.q.FindUserByDiscordID(ctx, discordID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        data.ID.String(),
		DiscordID: data.DiscordID,
		Pseudo:    data.Pseudo,
		Tag:       data.Tag,
	}, nil
}
