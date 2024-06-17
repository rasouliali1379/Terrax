package ports

import (
	"context"
	"github.com/rasouliali1379/terrax/internal/core/domain"
	"github.com/rasouliali1379/terrax/internal/core/domain/aggregate"
)

type ProfileRepository interface {
	Get(c context.Context, userID uint64) (*aggregate.Profile, error)
	New(c context.Context, user *domain.User) error
}
