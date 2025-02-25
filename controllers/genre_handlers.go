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

type GenresHandlers struct {
	DB *database.Queries
}

func NewGenresHandlers(db *database.Queries) *GenresHandlers {
	return &GenresHandlers{
		DB: db,
	}
}

// GetAll godoc
// @Tags Genres
// @Summary      Get Genres List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Success      200  {array} database.Genre "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get genres"
// @Router       /v1/genres [get]
// @Security Bearer
func (rh *GenresHandlers) GetAll(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Genres >= 2 {
			can_do = true
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	genres, err := rh.DB.GetGenres(r.Context())
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genres", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, genres)
}

// Create godoc
// @Tags Genres
// @Summary      Create Genre
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param request body views.CreateGenreRequest true "Genre data"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Create genre"
// @Router       /v1/genres [post]
// @Security Bearer
func (rh *GenresHandlers) Create(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Genres == 3 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	cgr := views.CreateGenreRequest{}

	err := decoder.Decode(&cgr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of CreateGenreRequest", err)
		return
	}

	id, err := rh.DB.CreateGenre(r.Context(), cgr.Title)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't create genre", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseId{
		ID: int(id),
	})
}

// Get godoc
// @Tags Genres
// @Summary      Get Genre
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} database.Genre "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get genre"
// @Router       /v1/roles/{id} [get]
// @Security Bearer
func (rh *GenresHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Genres >= 2 {
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

	genre, err := rh.DB.GetGenreById(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get genre", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, genre)
}

// Update godoc
// @Tags Genres
// @Summary      Update Genre
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.UpdateGenreRequest true "Genre data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Update genre"
// @Router       /v1/genres/{id} [put]
// @Security Bearer
func (gh *GenresHandlers) Update(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Genres == 3 {
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
	ugr := views.UpdateGenreRequest{}

	err = decoder.Decode(&ugr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of UpdateRoleRequest", err)
		return
	}

	err = gh.DB.UpdateGenre(r.Context(), database.UpdateGenreParams{
		ID:    int64(id),
		Title: ugr.Title,
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get roles", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Tags Genres
// @Summary      Delete Genre
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't delete genre"
// @Router       /v1/genres/{id} [delete]
// @Security Bearer
func (gh *GenresHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Genres >= 2 {
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

	err = gh.DB.DeleteGenre(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete roles", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseId{ID: id})
}
