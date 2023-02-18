package repository

import (
	"context"
	"examples/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID int) (*model.User, error)
	ModifyAuthority(ctx context.Context, userID, authority int) error
}
