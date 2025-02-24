package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
)

type RolesHandlers struct {
	DB *database.Queries
}

func NewRolesHandlers(db *database.Queries) *RolesHandlers {
	return &RolesHandlers{
		DB: db,
	}
}

// GetAll godoc
// @Tags Roles
// @Summary      Get Roles List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Success      200  {array} database.Role "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
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

// Get godoc
// @Tags Roles
// @Summary      Get Roles List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} database.Role "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get roles"
// @Router       /v1/roles/{id} [get]
// @Security Bearer
func (rh *RolesHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Roles >= 2 {
			can_do = true
			break
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

// Update godoc
// @Tags Roles
// @Summary      Get Roles List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.UpdateRoleRequest true "User data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Update role"
// @Router       /v1/roles/{id} [put]
// @Security Bearer
func (rh *RolesHandlers) Update(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Roles == 3 {
			can_do = true
			break
		}
	}

	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	urr := views.UpdateRoleRequest{}

	err = decoder.Decode(&urr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of UpdateRoleRequest", err)
		return
	}

	err = rh.DB.UpdateRole(r.Context(), database.UpdateRoleParams{
		ID:            int64(id),
		Title:         urr.Title,
		Projects:      urr.Projects,
		Genres:        urr.Genres,
		AgeCategories: urr.AgeCategories,
		Types:         urr.Types,
		Users:         urr.Users,
		Roles:         urr.Roles,
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get roles", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
