package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bayan2019/go-ozinshe/configuration"
	"github.com/Bayan2019/go-ozinshe/controllers"
	"github.com/Bayan2019/go-ozinshe/repositories"
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

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	// fmt.Println(dbURL)
	err := configuration.Connect2DB(dbURL)
	if err != nil {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
		fmt.Println(err.Error())
	}

	dir := os.Getenv("DIR")
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "superozinshe"
	}

	if configuration.ApiCfg != nil {
		configuration.ApiCfg.Dir = dir
		configuration.ApiCfg.JwtSecret = jwtSecret
	} else {
		fmt.Println("No DATABASE_URL")
		configuration.ApiCfg = &configuration.ApiConfiguration{
			Dir:       dir,
			JwtSecret: jwtSecret,
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

	router.Get("/", controllers.StaticHandler)

	router.Get("/hello", controllers.HelloHandler)

	router.Get("/swagger/*",
		httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/swagger/doc.json")))

	v1Router := chi.NewRouter()

	if configuration.ApiCfg.DB != nil {
		usersRepository := repositories.NewUsersRepository(configuration.ApiCfg.Conn)
		usersHandlers := controllers.NewUsersHandlers(usersRepository)

		v1Router.Post("/users", usersHandlers.Create)

		authHandlers := controllers.NewAuthHandlers(configuration.ApiCfg.DB, configuration.ApiCfg.JwtSecret)

		v1Router.Post("/auth/sign-in", authHandlers.Login)
		v1Router.Post("/auth/refresh", authHandlers.Refresh)
		v1Router.Post("/auth/sign-out", authHandlers.Logout)

		v1Router.Delete("/users", authHandlers.MiddlewareAuth(usersHandlers.Delete))

	}

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
		// ReadHeaderTimeout: time.Second * 5,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
