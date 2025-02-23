package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/Bayan2019/go-ozinshe/views"
)

type UsersHandlers struct {
	userRepo *repositories.UsersRepository
}

func NewUsersHandlers(repo *repositories.UsersRepository) *UsersHandlers {
	return &UsersHandlers{
		userRepo: repo,
	}
}

// Create godoc
// @Tags users
// @Summary      Create user
// @Accept       json
// @Produce      json
// @Param request body views.CreateUserRequest true "User data"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't hash password"
// @Router       /v1/users [post]
func (uh *UsersHandlers) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	cur := views.CreateUserRequest{}

	err := decoder.Decode(&cur)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of CreateUserRequest", err)
		return
	}

	hashedPassword, err := HashPassword(cur.Password)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	id, err := uh.userRepo.Create(r.Context(), repositories.CreateUserParams{
		Name:         cur.Name,
		Email:        cur.Email,
		PasswordHash: hashedPassword,
	})

	views.RespondWithJSON(w, http.StatusCreated, views.NewResponseId(int(id)))
}

// Delete godoc
// @Tags users
// @Summary      Delete user
// @Accept       json
// @Produce      json
// @Param id path int true "User id"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't delete user"
// @Router       /v1/users [delete]
// @Security Bearer
func (uh *UsersHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
	err := uh.userRepo.DB.DeleteUser(r.Context(), user.Id)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete user", err)
		return
	}
	views.RespondWithJSON(w, http.StatusOK, views.NewResponseId(int(user.Id)))
}
