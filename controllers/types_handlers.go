package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
)

type TypeHandlers struct {
	DB *database.Queries
}

func NewTypesHandlers(db *database.Queries) *TypeHandlers {
	return &TypeHandlers{
		DB: db,
	}
}

// GetAll godoc
// @Tags Types
// @Summary      Get Types List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Success      200  {array} database.Type "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get Types"
// @Router       /v1/types [get]
// @Security Bearer
func (th *TypeHandlers) GetAll(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Types >= 2 {
			can_do = true
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	types, err := th.DB.GetTypes(r.Context())
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genres", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, types)
}

// Create godoc
// @Tags Types
// @Summary      Create Type
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param request body views.CreateTypeRequest true "Type data"
// @Success      201  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Create Type"
// @Router       /v1/types [post]
// @Security Bearer
func (th *TypeHandlers) Create(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Types == 3 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	ctr := views.CreateTypeRequest{}

	err := decoder.Decode(&ctr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of CreateTypeRequest", err)
		return
	}

	id, err := th.DB.CreateType(r.Context(), ctr.Title)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't create Type", err)
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, views.ResponseId{
		ID: int(id),
	})
}

// Get godoc
// @Tags Types
// @Summary      Get Type
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} database.Type "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get Type"
// @Router       /v1/types/{id} [get]
// @Security Bearer
func (th *TypeHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Types >= 2 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	type1, err := th.DB.GetTypeById(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genre", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, type1)
}

// Update godoc
// @Tags Types
// @Summary      Update Type
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.UpdateTypeRequest true "Type data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Update Type"
// @Router       /v1/types/{id} [put]
// @Security Bearer
func (th *TypeHandlers) Update(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Types == 3 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	utr := views.UpdateTypeRequest{}

	err = decoder.Decode(&utr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of UpdateTypeRequest", err)
		return
	}

	err = th.DB.UpdateType(r.Context(), database.UpdateTypeParams{
		ID:    int64(id),
		Title: utr.Title,
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't Update Type", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Tags Types
// @Summary      Delete Type
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't delete Type"
// @Router       /v1/types/{id} [delete]
// @Security Bearer
func (th *TypeHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Types >= 2 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	err = th.DB.DeleteType(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete roles", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseId{ID: id})
}
