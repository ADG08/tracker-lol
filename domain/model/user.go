package model

import "github.com/google/uuid"

type User struct {
	ID        string
	DiscordID string
	Pseudo    string
	Tag       string
}

func CreateUser(discordID, pseudo, tag string) *User {
	return &User{
		ID:        uuid.New().String(),
		DiscordID: discordID,
		Pseudo:    pseudo,
		Tag:       tag,
	}
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetDiscordID() string {
	return u.DiscordID
}

func (u *User) GetPseudo() string {
	return u.Pseudo
}

func (u *User) GetTag() string {
	return u.Tag
}
