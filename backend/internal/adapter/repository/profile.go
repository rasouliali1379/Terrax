package repository

import (
	"context"
	"errors"
	"github.com/rasouliali1379/terrax/internal/core/domain"
	"github.com/rasouliali1379/terrax/internal/core/domain/aggregate"
	"github.com/rasouliali1379/terrax/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository struct {
	db *mongo.Database
}

func NewProfileRepo(db *mongo.Database) ports.ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (u ProfileRepository) Get(c context.Context, userID uint64) (*aggregate.Profile, error) {
	filter := bson.M{"user.user_id": userID}

	var profile aggregate.Profile
	if err := u.db.Collection("profiles").FindOne(c, filter).Decode(&profile); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (u ProfileRepository) New(c context.Context, user *domain.User) error {
	if _, err := u.db.
		Collection("profiles").
		InsertOne(c, aggregate.Profile{User: user, Tap: &domain.Tap{Count: 0, Energy: 1000}}); err != nil {
		return err
	}
	return nil
}
