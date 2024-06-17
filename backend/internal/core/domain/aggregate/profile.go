package aggregate

import "github.com/rasouliali1379/terrax/internal/core/domain"

type Profile struct {
	User *domain.User
	Tap  *domain.Tap
}
