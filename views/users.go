package views

import (
	"github.com/Bayan2019/go-ozinshe/repositories/database"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	// Id          int64     `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Phone       string `json:"phone"`
}

type UpdateUserRequest struct {
	// Id          int64     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	DateOfBirth string  `json:"date_of_birth"`
	Phone       string  `json:"phone"`
	RoleIds     []int64 `json:"role_ids"`
}

type User struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	DateOfBirth string          `json:"date_of_birth"`
	Phone       string          `json:"phone"`
	Roles       []database.Role `json:"roles"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
