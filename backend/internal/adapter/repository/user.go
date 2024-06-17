package repository

import (
	"github.com/rasouliali1379/terrax/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) ports.UserRepository {
	return UserRepository{
		db: db,
	}
}
