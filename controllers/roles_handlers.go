package controllers

import (
	"errors"
	"net/http"

	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/Bayan2019/go-ozinshe/views"
)

type RolesHandlers struct {
	DB *repositories.Queries
}

func NewRolesHandlers(db *repositories.Queries) *RolesHandlers {
	return &RolesHandlers{
		DB: db,
	}
}

// Delete godoc
// @Tags Roles
// @Summary      Get Roles List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "With the Bearer started"
// @Success      200  {array} repositories.Role "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get roles"
// @Router       /v1/roles [get]
// @Security Bearer
func (rh *RolesHandlers) GetAll(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Roles >= 2 {
			can_do = true
		}
	}

	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	roles, err := rh.DB.GetRoles(r.Context())
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get roles", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, roles)
}
