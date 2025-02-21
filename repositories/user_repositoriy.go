package repositories

import "database/sql"

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
