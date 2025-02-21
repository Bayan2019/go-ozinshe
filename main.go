package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Bayan2019/go-ozinshe/configuration"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	// ginSwagger "github.com/swaggo/gin-swagger"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/Bayan2019/go-ozinshe/docs"
	// _ "github.com/mattn/go-sqlite3"
)

//go:embed static/*
var staticFiles embed.FS

// @title ÖZINŞE API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		platform = "dev"
	}

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
	err := configuration.Connect2DB(dbURL, platform)
	if err != nil {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
		fmt.Println(err.Error())
	}

	dirImages := os.Getenv("DIR_IMAGES")
	if dirImages == "" {
		dirImages = "/images"
	}

	dirVideos := os.Getenv("DIR_VIDEOS")
	if dirVideos == "" {
		dirVideos = "/videos"
	}

	if configuration.ApiCfg != nil {
		configuration.ApiCfg.DirImages = dirImages
		configuration.ApiCfg.DirVideos = dirVideos
	} else {
		fmt.Println("No DATABASE_URL")
		configuration.ApiCfg = &configuration.ApiConfiguration{
			DirImages: dirImages,
			DirVideos: dirVideos,
		}
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if _, err := io.Copy(w, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.Get("/api", func(w http.ResponseWriter, r *http.Request) {

		views.RespondWithJSON(w, http.StatusOK, views.ResponseMessage{
			Message: "hello from api",
		})
	})

	router.Get("/swagger/*",
		httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/swagger/doc.json")))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
		// ReadHeaderTimeout: time.Second * 5,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
