package controllers

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type ImagesHandlers struct {
	DB  *database.Queries
	Dir string
}

func NewImagesHandlers(db *database.Queries, dir string) *ImagesHandlers {
	return &ImagesHandlers{
		DB:  db,
		Dir: dir,
	}
}

// Display godoc
// @Tags Images
// @Summary      Display Image
// @Accept       json
// @Produce      application/octet-stream
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "No Permission" "Couldn't get User Middleware"
// @Router       /v1/projects/images/show/{id} [get]
// @Security Bearer
func (ih *ImagesHandlers) Display(w http.ResponseWriter, r *http.Request, user views.User) {
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
	// You can get the string value of the path parameter like in Go
	// with the http.Request.PathValue method.
	id := chi.URLParam(r, "id")
	// enableCors(&w)
	w.Header().Set("Content-Type", "image/jpeg")
	// w.ResponseWriter.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, fmt.Sprintf("%s%s", ih.Dir, id))
}

// Display godoc
// @Tags Images
// @Summary      Get Image
// @Accept       json
// @Produce      application/octet-stream
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "can't read the image"
// @Router       /v1/projects/images [get]
// @Security Bearer
func (ih *ImagesHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
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
	id := chi.URLParam(r, "id")
	byteFile, err := os.ReadFile(fmt.Sprintf("%s%s", ih.Dir, id))
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "can't read the image", err)
		return
	}
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", mediaTypeToExt(id)))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", id))
	w.WriteHeader(http.StatusOK)
	w.Write(byteFile)
}

// Display godoc
// @Tags Images
// @Summary      Create Images
// @Accept       multipart/form-data
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param poster_id formData int true "poster_id"
// @Param poster formData file true "image"
// @Success      200  {object} views.ResponseMessage  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "can't create image"
// @Router       /v1/projects/images [post]
// @Security Bearer
func (ih *ImagesHandlers) Upload(w http.ResponseWriter, r *http.Request, user views.User) {
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
	// Use r.FormFile to get the file data. The key the web browser is using is called "thumbnail"
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
	fpath := fmt.Sprintf("%s%s", ih.Dir, fileName)
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

	project_id, err := strconv.Atoi(r.FormValue("project_id"))

	err = ih.DB.AddImage2Movie(r.Context(), database.AddImage2MovieParams{
		ID:        fileName,
		ProjectID: int64(project_id),
	})
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error saving file", err)
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, views.ResponseMessage{
		Message: fileName,
	})
}

// Display godoc
// @Tags Images
// @Summary      Delete Images
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  {object} views.ResponseMessage  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "Error deleting file"
// @Router       /v1/projects/images [delete]
// @Security Bearer
func (ih *ImagesHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
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

	id := chi.URLParam(r, "id")
	fpath := fmt.Sprintf("%s%s", ih.Dir, id)
	// Use os.Create to create the new file
	err := os.Remove(fpath)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error deleting file", err)
		return
	}

	err = ih.DB.DeleteImage(r.Context(), id)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error deleting file", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseMessage{
		Message: id,
	})
}

func mediaTypeToExt(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}
