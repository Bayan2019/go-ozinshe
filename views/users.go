package views

import (
	"time"

	"github.com/Bayan2019/go-ozinshe/repositories"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id          int64               `json:"id"`
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	DateOfBirth time.Time           `json:"date_of_birth"`
	Phone       string              `json:"phone"`
	Roles       []repositories.Role `json:"roles"`
}
