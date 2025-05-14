package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
}

type DiscordConfig struct {
	DiscordToken string `env:"DISCORD_TOKEN,required"`
	ChannelID    string `env:"CHANNEL_ID,required"`
}

func Load() (*DiscordConfig, *DatabaseConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, err
	}

	discordCfg := &DiscordConfig{}
	if err := env.Parse(discordCfg); err != nil {
		return nil, nil, err
	}

	dbCfg := &DatabaseConfig{}
	if err := env.Parse(dbCfg); err != nil {
		return nil, nil, err
	}

	return discordCfg, dbCfg, nil
}
