package repository

import (
	"context"
	"examples/code"
	"examples/errors"
	"examples/pkg/http/adapter/infra"
	"examples/pkg/http/adapter/repository/model"
	"examples/pkg/http/entity"
	"examples/pkg/http/logic/repository"
	"time"
)

type loginRepository struct {
	infra.SqlHandler
}

func NewLoginRepository(handler infra.SqlHandler) *loginRepository {
	return &loginRepository{handler}
}

func (r *loginRepository) GetByID(ctx context.Context, loginID string) (*entity.Login, error) {
	query := "SELECT * FROM login WHERE login_id = ?"
	row, err := r.QueryRow(ctx, query, loginID)
	if err != nil {
		return nil, err
	}

	var login model.Login
	if err := row.Scan(&login.LoginID, &login.UserID, &login.LastSignedAt, &login.Password); err != nil {
		return nil, errors.Wrap(code.ErrNotFound, err)
	}
	return &entity.Login{
		LoginID:      login.LoginID,
		UserID:       login.UserID,
		LastSignedAt: login.LastSignedAt,
		Password:     login.Password,
	}, nil
}

func (r *loginRepository) GetByUserID(ctx context.Context, userID int) (*entity.Login, error) {
	query := "SELECT * FROM login WHERE user_id = ?"
	row, err := r.QueryRow(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var login model.Login
	if err := row.Scan(&login.LoginID, &login.UserID, &login.LastSignedAt, &login.Password); err != nil {
		return nil, errors.Wrap(code.ErrNotFound, err)
	}
	return &entity.Login{
		LoginID:      login.LoginID,
		UserID:       login.UserID,
		LastSignedAt: login.LastSignedAt,
		Password:     login.Password,
	}, nil
}

func (r *loginRepository) Store(ctx context.Context, login *entity.Login) error {
	query := "INSERT INTO login (`login_id`, `user_id`, `password`) VALUES (?, ?, ?)"
	if _, err := r.Execute(ctx, query, login.LoginID, login.UserID, login.Password); err != nil {
		return err
	}
	return nil
}

func (r *loginRepository) ModifyLastSigned(ctx context.Context, loginID string) error {
	query := "UPDATE login SET last_signed_at = strftime('%s', ?) WHERE login_id = ?"
	if _, err := r.Execute(ctx, query, time.Now(), loginID); err != nil {
		return err
	}
	return nil
}

var _ repository.LoginRepository = (*loginRepository)(nil)