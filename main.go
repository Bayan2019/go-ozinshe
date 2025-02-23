package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bayan2019/go-ozinshe/configuration"
	"github.com/Bayan2019/go-ozinshe/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	// ginSwagger "github.com/swaggo/gin-swagger"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/Bayan2019/go-ozinshe/docs"
	// _ "github.com/mattn/go-sqlite3"
)

// @title ÖZINŞE API
// @version 1.0
// @description This is a sample server ÖZINŞE.
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

	// platform := os.Getenv("PLATFORM")
	// if platform == "" {
	// 	platform = "dev"
	// }

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
	err := configuration.Connect2DB(dbURL)
	if err != nil {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
		fmt.Println(err.Error())
	}

	dir := os.Getenv("DIR")

	if configuration.ApiCfg != nil {
		configuration.ApiCfg.Dir = dir
	} else {
		fmt.Println("No DATABASE_URL")
		configuration.ApiCfg = &configuration.ApiConfiguration{}
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

	router.Get("/", controllers.StaticHandler)

	router.Get("/hello", controllers.HelloHandler)

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
