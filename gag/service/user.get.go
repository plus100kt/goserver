package service

import (
	"context"

	"github.com/plus100kt/goserver/gag/model"
)

// Get retrieves a user based on their uuid
func (s *userService) Get(ctx context.Context, id string) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, id)

	return u, err
}
