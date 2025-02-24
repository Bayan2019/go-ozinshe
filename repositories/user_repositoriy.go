package repositories

import (
	"context"
	"database/sql"

	"github.com/Bayan2019/go-ozinshe/views"
)

type UsersRepository struct {
	Conn *sql.DB
	DB   *Queries
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		Conn: db,
		DB:   New(db),
	}
}

func (ur *UsersRepository) Create(ctx context.Context, cup CreateUserParams) (int64, error) {
	tx, err := ur.Conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	qtx := ur.DB.WithTx(tx)

	id, err := qtx.CreateUser(ctx, cup)
	if err != nil {
		return 0, err
	}

	err = qtx.AddRole2User(ctx, AddRole2UserParams{
		UserID: id,
		RoleID: 2,
	})
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (ur *UsersRepository) UpdateProfile(ctx context.Context, upr views.UpdateProfileRequest) error {
	tx, err := ur.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := ur.DB.WithTx(tx)

	err = qtx.UpdateUser(ctx, UpdateUserParams{
		ID:          upr.Id,
		Name:        upr.Name,
		Email:       upr.Email,
		DateOfBirth: upr.DateOfBirth,
		Phone:       upr.Phone,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ur *UsersRepository) Delete(ctx context.Context, id int64) error {
	tx, err := ur.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := ur.DB.WithTx(tx)

	err = qtx.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
