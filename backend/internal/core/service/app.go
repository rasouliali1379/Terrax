package service

import (
	"context"
	"github.com/rasouliali1379/terrax/internal/core/domain"
	"github.com/rasouliali1379/terrax/internal/core/domain/aggregate"
	"github.com/rasouliali1379/terrax/internal/core/ports"
)

type AppService struct {
	profileRepo ports.ProfileRepository
}

func NewAppService(profileRepo ports.ProfileRepository) ports.AppService {
	return &AppService{profileRepo: profileRepo}
}

func (a AppService) GetUserData(c context.Context, user *domain.User) (*aggregate.Profile, error) {
	profile, err := a.profileRepo.Get(c, user.UserID)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		if err := a.profileRepo.New(c, user); err != nil {
			return nil, err
		}

		profile, err = a.profileRepo.Get(c, user.UserID)
		if err != nil {
			return nil, err
		}
	}

	return profile, nil
}
