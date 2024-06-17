package repository

import (
	"github.com/rasouliali1379/terrax/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type TapRepository struct {
	db *mongo.Database
}

func NewTapRepo(db *mongo.Database) ports.TapRepository {
	return TapRepository{
		db: db,
	}
}
