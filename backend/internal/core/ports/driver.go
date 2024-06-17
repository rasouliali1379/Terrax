package ports

import (
	"context"
	"github.com/rasouliali1379/terrax/internal/core/domain"
	"github.com/rasouliali1379/terrax/internal/core/domain/aggregate"
)

type AppService interface {
	GetUserData(c context.Context, user *domain.User) (*aggregate.Profile, error)
}
