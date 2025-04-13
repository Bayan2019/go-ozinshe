package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type ProjectsHandlers struct {
	repo *repositories.ProjectsRepository
	Dir  string
}

func NewProjecsHandlers(repo *repositories.ProjectsRepository, dir string) *ProjectsHandlers {
	return &ProjectsHandlers{
		repo: repo,
		Dir:  dir,
	}
}

// GetAll godoc
// @Tags Projects
// @Summary      Get Projects List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Success      200  {array} views.Project "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get Projects"
// @Router       /v1/projects [get]
// @Security Bearer
func (ph *ProjectsHandlers) GetAll(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Projects >= 2 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	// idsArray := r.URL.Query()["genre_id"]
	// fmt.Println(idsArray)

	projects, err := ph.repo.GetAll(r.Context())
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get projects", err)
	}

	views.RespondWithJSON(w, http.StatusOK, projects)
}

// GetAllSearchTerm godoc
// @Tags Projects
// @Summary      Get Projects List
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param searchTerm query string false "Search Term"
// @Param genre_id query []string false "Genre Ids" collectionFormat(multi)
// @Success      200  {array} views.Project "OK"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get Projects"
// @Router       /v1/projects/search [get]
// @Security Bearer
func (ph *ProjectsHandlers) GetAllSearch(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Projects >= 2 {
			can_do = true
			break
		}
	}
	if !can_do {
		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
		return
	}

	// idsArray := r.URL.Query()["genre_id"]
	// fmt.Println(idsArray)
	searchTerm := r.URL.Query().Get("searchTerm")
	idsArray := r.URL.Query()["genre_id"]

	if searchTerm == "" && len(idsArray) == 0 {
		views.RespondWithJSON(w, http.StatusOK, []views.Project{})
		return
	}

	ids := []int64{}
	for _, genre_id := range idsArray {
		id, err := strconv.Atoi(genre_id)
		if err != nil {
			views.RespondWithError(w, http.StatusBadRequest, "wrong genre_id", err)
			return
		}
		ids = append(ids, int64(id))
	}

	if searchTerm == "" && len(ids) != 0 {
		dProjects, err := ph.repo.DB.GetProjectsOfGenrers(r.Context(), ids)
		if err != nil {
			views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get projects of genres", err)
			return
		}
		projects, err := ph.repo.DatabaseProjects2viewsProjects(r.Context(), dProjects)
		if err != nil {
			views.RespondWithError(w, http.StatusInternalServerError, "Couldn't convert database projects to views projects", err)
			return
		}
		views.RespondWithJSON(w, http.StatusOK, projects)
		return
	}

	if searchTerm != "" && len(idsArray) == 0 {
		// err = ph.repo.DB.
		searchTerm = "%" + strings.ToLower(searchTerm) + "%"
		dProjects, err := ph.repo.DB.GetProjectsSearch(r.Context(), searchTerm)
		if err != nil {
			views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get projects of search term", err)
			return
		}
		projects, err := ph.repo.DatabaseProjects2viewsProjects(r.Context(), dProjects)
		if err != nil {
			views.RespondWithError(w, http.StatusInternalServerError, "Couldn't convert database projects to views projects", err)
			return
		}
		views.RespondWithJSON(w, http.StatusOK, projects)
		return
	}

	// var projects []views.RProject

	dProjects, err := ph.repo.DB.GetProjectsOfGenresAndSearch(r.Context(), database.GetProjectsOfGenresAndSearchParams{
		Ids:    ids,
		Search: searchTerm,
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get projects of search term and genres", err)
		return
	}

	projects, err := ph.repo.DatabaseProjects2viewsProjects(r.Context(), dProjects)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't convert database projects to views projects", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, projects)
}

// GetProject godoc
// @Tags Projects
// @Summary      Get Project
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} views.Project "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Get Project"
// @Router       /v1/projects/{id} [get]
// @Security Bearer
func (ph *ProjectsHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
	can_do := false
	for _, role := range user.Roles {
		if role.Projects >= 2 {
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

	project, err := ph.repo.GetById(r.Context(), int64(id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't get Project", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, project)
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

// Update godoc
// @Tags Projects
// @Summary      Set Cover for Project
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param request body views.ImageIdRequest true "Project data"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Couldn't Set Cover for Project"
// @Router       /v1/projects/{id}/cover [patch]
// @Security Bearer
func (ph *ProjectsHandlers) SetCover(w http.ResponseWriter, r *http.Request, user views.User) {
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

	project_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	imageIdR := views.ImageIdRequest{}

	err = decoder.Decode(&imageIdR)
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Error parsing JSON of ImageIdRequest", err)
		return
	}

	err = ph.repo.DB.SetCover(r.Context(), database.SetCoverParams{
		Cover: sql.NullString{
			String: imageIdR.ImageId,
			Valid:  true,
		},
		ID: int64(project_id),
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Couldn't set the Cover for the Project", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Upload godoc
// @Tags Projects
// @Summary      Upload Cover
// @Accept       multipart/form-data
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Param image formData file true "image"
// @Success      200  {object} views.ResponseIdStr  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "can't create image"
// @Router       /v1/projects/{id}/cover [post]
// @Security Bearer
func (ph *ProjectsHandlers) UploadCover(w http.ResponseWriter, r *http.Request, user views.User) {
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

	// Set a const maxMemory to 10MB.
	const maxMemory = 10 << 20 // 10 MB
	// Use (http.Request).ParseMultipartForm with the maxMemory const as an argument
	r.ParseMultipartForm(maxMemory)
	// Use r.FormFile to get the file data. The key the web browser is using is called "image"
	file, header, err := r.FormFile("image")
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Unable to parse form file", err)
		return
	}
	defer file.Close()

	// Get the media type from the file's Content-Type header
	// Use the mime.ParseMediaType function to get the media type from the Content-Type header
	mediaType, _, err := mime.ParseMediaType(header.Header.Get("Content-Type"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid Content-Type", err)
		return
	}
	// If the media type isn't either image/jpeg or image/png,
	// respond with an error (respondWithError helper)
	if mediaType != "image/jpeg" && mediaType != "image/png" {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid file type", nil)
		return
	}
	ext := mediaTypeToExt(mediaType)
	fileName := fmt.Sprintf("%s%s", uuid.NewString(), ext)
	fpath := fmt.Sprintf("%s%s", ph.Dir, fileName)
	// Use os.Create to create the new file
	dst, err := os.Create(fpath)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Unable to create file on server", err)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
		return
	}

	project_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}

	err = ph.repo.UploadCover(r.Context(), int64(project_id), fileName)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, views.ResponseIdStr{
		ID: fileName,
	})
}
