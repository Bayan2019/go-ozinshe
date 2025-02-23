package controllers

import (
	"embed"
	"io"
	"net/http"

	"github.com/Bayan2019/go-ozinshe/views"
)

//go:embed static/*
var staticFiles embed.FS

// Static godoc
// @Summary      Giving Common page
// @Tags Tests
// @Produce      html
// @Success      200  {body} file "OK"
// @Failure   	 500  {object} views.ErrorResponse "Invalid file"
// @Router       / [get]
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	f, err := staticFiles.Open("static/index.html")
	if err != nil {
		views.RespondWithError(w, http.StatusInsufficientStorage, "can't open static/index.html", err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Static godoc
// @Summary      Saying hello
// @Tags Tests
// @Produce      json
// @Success      200  {object} views.ResponseMessage "OK"
// @Failure   	 500  {object} views.ErrorResponse "Invalid file"
// @Router       /hello [get]
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	views.RespondWithJSON(w, http.StatusOK, views.ResponseMessage{
		Message: "hello from api",
	})
}
