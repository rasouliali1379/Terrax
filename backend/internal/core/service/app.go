package service

import "github.com/rasouliali1379/terrax/internal/core/ports"

type AppService struct {
}

func NewAppService() ports.AppService {
	return &AppService{}
}
