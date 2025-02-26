package controllers

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strconv"

	"github.com/Bayan2019/go-ozinshe/repositories/database"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type VideosHandlers struct {
	DB  *database.Queries
	Dir string
}

func NewVideosHandlers(db *database.Queries, dir string) *VideosHandlers {
	return &VideosHandlers{
		DB:  db,
		Dir: dir,
	}
}

// Display godoc
// @Tags Videos
// @Summary      Get Video
// @Accept       json
// @Produce      application/octet-stream
// @Param Authorization header string true "Bearer AccessToken"
// @Param id path int true "id"
// @Success      200  "OK"
// @Failure   	 400  {object} views.ErrorResponse "Invalid data"
// @Failure   	 401  {object} views.ErrorResponse "No token Middleware"
// @Failure   	 403  {object} views.ErrorResponse "No Permission"
// @Failure   	 404  {object} views.ErrorResponse "Not found User Middleware"
// @Failure   	 500  {object} views.ErrorResponse "can't read the video"
// @Router       /v1/projects/videos [get]
// @Security Bearer
// func (vh *VideosHandlers) Get(w http.ResponseWriter, r *http.Request, user views.User) {
// 	can_do := false
// 	for _, role := range user.Roles {
// 		if role.Projects >= 2 {
// 			can_do = true
// 			break
// 		}
// 	}
// 	if !can_do {
// 		views.RespondWithError(w, http.StatusForbidden, "don't have permission", errors.New("no Permission"))
// 		return
// 	}

// 	id := chi.URLParam(r, "id")
// 	byteFile, err := os.ReadFile(fmt.Sprintf("%s%s", vh.Dir, id))
// 	if err != nil {
// 		views.RespondWithError(w, http.StatusInternalServerError, "can't read the video", err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "video/mp4")
// 	w.Header().Set("Content-Type", "application/octet-stream")
// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", id))
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(byteFile)
// }

// Display godoc
// @Tags Videos
// @Summary      Create Video
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
// @Router       /v1/projects/videos [post]
// @Security Bearer
func (vh *VideosHandlers) Upload(w http.ResponseWriter, r *http.Request, user views.User) {
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

	// Set an upload limit of 1 GB (1 << 30 bytes)
	const uploadLimit = 1 << 26
	// using http.MaxBytesReader
	r.Body = http.MaxBytesReader(w, r.Body, uploadLimit)

	// Use r.FormFile to get the file data. The key the web browser is using is called "image"
	file, header, err := r.FormFile("video")
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Unable to parse form file", err)
		return
	}
	defer file.Close()

	// Validate the uploaded file to ensure it's an MP4 video
	// Use mime.ParseMediaType and "video/mp4" as the MIME type
	mediaType, _, err := mime.ParseMediaType(header.Header.Get("Content-Type"))
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid Content-Type", err)
		return
	}
	if mediaType != "video/mp4" {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid file type, only MP4 is allowed", nil)
		return
	}

	fileName := fmt.Sprintf("%s.mp4", uuid.NewString())
	fpath := fmt.Sprintf("%s%s", vh.Dir, fileName)
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
	if err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid project_id", err)
		return
	}

	err = vh.DB.AddVideo2Movie(r.Context(), database.AddVideo2MovieParams{
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
// @Tags Videos
// @Summary      Delete Video
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
// @Router       /v1/projects/videos [delete]
// @Security Bearer
func (vh *VideosHandlers) Delete(w http.ResponseWriter, r *http.Request, user views.User) {
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
	fpath := fmt.Sprintf("%s%s", vh.Dir, id)
	// Use os.Create to create the new file
	err := os.Remove(fpath)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error deleting file", err)
		return
	}

	err = vh.DB.DeleteVideo(r.Context(), id)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Error deleting file", err)
		return
	}

	views.RespondWithJSON(w, http.StatusOK, views.ResponseMessage{
		Message: id,
	})
}
