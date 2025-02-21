package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Bayan2019/go-ozinshe/configuration"
	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/joho/godotenv"
)

// @title           Ozinshe API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      ozinshe.sapar.com
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dirImages := os.Getenv("DIR_IMAGES")
	if dirImages == "" {
		dirImages = "/images"
	}
	dirVideos := os.Getenv("DIR_VIDEOS")
	if dirVideos == "" {
		dirVideos = "/videos"
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		platform = "dev"
	}

	configuration.ApiCfg = &configuration.ApiConfiguration{
		DirImages: dirImages,
		DirVideos: dirVideos,
	}

	// https://github.com/libsql/libsql-client-go/#open-a-connection-to-sqld
	// libsql://[your-database].turso.io?authToken=[your-auth-token]
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
	} else {
		var db repositories.DBTX

		if platform == "dev" {
			db1, err := sql.Open("sqlite3", dbURL)
			if err != nil {
				log.Fatal(err)
			}
			db = db1
		} else {
			db1, err := sql.Open("libsql", dbURL)
			if err != nil {
				log.Fatal(err)
			}
			db = db1
		}

		dbQueries := repositories.New(db)
		configuration.ApiCfg.DB = dbQueries
		log.Println("Connected to database!")
	}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())
}
