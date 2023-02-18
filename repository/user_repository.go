package repository

import (
	"context"
	"examples/model"
)

type userRepository struct {
	SqlHandler
}

func NewUserRepository(handler SqlHandler) *userRepository {
	return &userRepository{handler}
}

func (r *userRepository) GetByID(ctx context.Context, userID int) (*model.User, error) {
	query := "SELECT * FROM user WHERE user_id = ?"
	row, err := r.QueryRow(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var user *model.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Authority, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ModifyAuthority(ctx context.Context, userID, authority int) error {
	query := "UPDATE user SET authority = ? WHERE user_id = ?"
	if _, err := r.Execute(ctx, query, authority, userID); err != nil {
		return err
	}
	return nil
}
