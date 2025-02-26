package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
)

type ProjectsHandlers struct {
	repo *repositories.ProjectsRepository
}

func NewProjecsHandlers(repo *repositories.ProjectsRepository) *ProjectsHandlers {
	return &ProjectsHandlers{
		repo: repo,
	}
}

// Create godoc
// @Tags Projects
// @Summary      Create Project
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param request body views.CreateProjectRequest true "Project data"
// @Success      201  {object} views.ResponseId "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Create Project"
// @Router       /v1/projects [post]
// @Security Bearer
func (ph *ProjectsHandlers) Create(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Projects == 3 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	cpr := views.CreateProjectRequest{}

	err := decoder.Decode(&cpr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of CreateProjectRequest", err)
		return
	}

	id, err := ph.repo.Create(r.Context(), cpr)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't create project", err)
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, views.ResponseId{
		ID: int(id),
	})
}

// Update godoc
// @Tags Projects
// @Summary      Update Project
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.UpdateProjectRequest true "Project data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Update Project"
// @Router       /v1/projects/{id} [put]
// @Security Bearer
func (ph *ProjectsHandlers) Update(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Projects == 3 {
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
	upr := views.UpdateProjectRequest{}

	err = decoder.Decode(&upr)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of UpdateProjectRequest", err)
		return
	}

	err = ph.repo.Update(r.Context(), int64(id), upr)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't Update Project", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
