package config

import (
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/rs/zerolog/log"
)

func ConnectDB() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to database")
	return client, nil
}
