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

type AgeCategoriesHandlers struct {
	DB *database.Queries
}

func NewAgeCategoriesHandlers(db *database.Queries) *AgeCategoriesHandlers {
	return &AgeCategoriesHandlers{
		DB: db,
	}
}

// GetAll godoc
// @Tags AgeCategories
// @Summary      Get AgeCategories List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Success      200  {array} database.AgeCategory "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get AgeCategories"
// @Router       /v1/age-categories [get]
// @Security Bearer
func (ach *AgeCategoriesHandlers) GetAll(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.AgeCategories >= 2 {
			can_do = true
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	genres, err := ach.DB.GetAgeCategories(r.Context())
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genres", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, genres)
}

// Create godoc
// @Tags AgeCategories
// @Summary      Create AgeCategory
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param request body views.CreateAgeCategoryRequest true "AgeCategory data"
// @Success      201  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Create AgeCategory"
// @Router       /v1/age-categories [post]
// @Security Bearer
func (ach *AgeCategoriesHandlers) Create(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.AgeCategories == 3 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	cacr := views.CreateAgeCategoryRequest{}

	err := decoder.Decode(&cacr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of CreateAgeCategoryRequest", err)
		return
	}

	id, err := ach.DB.CreateAgeCategory(r.Context(), cacr.Title)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't create genre", err)
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, views.ResponseId{
		ID: int(id),
	})
}

// Get godoc
// @Tags AgeCategories
// @Summary      Get AgeCategory
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} database.AgeCategory "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get AgeCategory"
// @Router       /v1/age-categories/{id} [get]
// @Security Bearer
func (ach *AgeCategoriesHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.AgeCategories >= 2 {
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

	age_category, err := ach.DB.GetAgeCategoryById(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genre", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, age_category)
}

// Update godoc
// @Tags AgeCategories
// @Summary      Update AgeCategory
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.UpdateAgeCategoryRequest true "AgeCategory data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Update AgeCategory"
// @Router       /v1/age-categories/{id} [put]
// @Security Bearer
func (ach *AgeCategoriesHandlers) Update(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.AgeCategories == 3 {
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
	uacr := views.UpdateAgeCategoryRequest{}

	err = decoder.Decode(&uacr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of UpdateAgeCategoryRequest", err)
		return
	}

	err = ach.DB.UpdateAgeCategory(r.Context(), database.UpdateAgeCategoryParams{
		ID:    int64(id),
		Title: uacr.Title,
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't update AgeCategory", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Tags AgeCategories
// @Summary      Delete AgeCategory
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't delete AgeCategory"
// @Router       /v1/age-categories/{id} [delete]
// @Security Bearer
func (ach *AgeCategoriesHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.AgeCategories >= 2 {
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

	err = ach.DB.DeleteAgeCategory(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete roles", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseId{ID: id})
}
