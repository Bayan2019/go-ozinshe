// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
)

const changePassword = `-- name: ChangePassword :exec

UPDATE users
SET password_hash = ?
WHERE id = ?
`

type ChangePasswordParams struct {
	PasswordHash string
	ID           int64
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) error {
	_, err := q.db.ExecContext(ctx, changePassword, arg.PasswordHash, arg.ID)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users(name, email, password_hash)
VALUES (?, ?, ?)
RETURNING id
`

type CreateUserParams struct {
	Name         string
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.PasswordHash)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec

DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one

SELECT id, created_at, updated_at, name, email, password_hash, date_of_birth, phone FROM users WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.DateOfBirth,
		&i.Phone,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one

SELECT id, created_at, updated_at, name, email, password_hash, date_of_birth, phone FROM users WHERE id = ?
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.DateOfBirth,
		&i.Phone,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many

SELECT id, created_at, updated_at, name, email, password_hash, date_of_birth, phone FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Email,
			&i.PasswordHash,
			&i.DateOfBirth,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersOfRole = `-- name: GetUsersOfRole :many

SELECT u.id, u.created_at, u.updated_at, u.name, u.email, u.password_hash, u.date_of_birth, u.phone
FROM users AS u
JOIN users_roles AS ur
ON u.id = ur.user_id
WHERE ur.role_id = ?
`

func (q *Queries) GetUsersOfRole(ctx context.Context, roleID int64) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsersOfRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Email,
			&i.PasswordHash,
			&i.DateOfBirth,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec

UPDATE users
SET updated_at = CURRENT_TIMESTAMP,
    name = ?,
    email = ?,
    date_of_birth = ?,
    phone = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name        string
	Email       string
	DateOfBirth string
	Phone       string
	ID          int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.DateOfBirth,
		arg.Phone,
		arg.ID,
	)
	return err
}
