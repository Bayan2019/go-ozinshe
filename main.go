package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("warning: assuming default configuration. .env unreadable: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	// fmt.Println(dbURL)
	err = configuration.Connect2DB(dbURL)
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
		authHandlers := controllers.NewAuthHandlers(configuration.ApiCfg.DB, configuration.ApiCfg.JwtSecret)

		v1Router.Post("/auth/sign-in", authHandlers.Login)
		v1Router.Post("/auth/refresh", authHandlers.Refresh)
		v1Router.Post("/auth/sign-out", authHandlers.Logout)

		usersRepository := repositories.NewUsersRepository(configuration.ApiCfg.Conn)
		usersHandlers := controllers.NewUsersHandlers(usersRepository)

		v1Router.Post("/users", usersHandlers.Register)
		v1Router.Get("/users", authHandlers.MiddlewareAuth(usersHandlers.GetUsers))
		v1Router.Get("/users/{id}", authHandlers.MiddlewareAuth(usersHandlers.GetUser))
		v1Router.Put("/users/{id}", authHandlers.MiddlewareAuth(usersHandlers.Update))
		v1Router.Delete("/users/{id}", authHandlers.MiddlewareAuth(usersHandlers.Delete))

		v1Router.Get("/users/profile", authHandlers.MiddlewareAuth(usersHandlers.GetProfile))
		v1Router.Put("/users/profile", authHandlers.MiddlewareAuth(usersHandlers.UpdateProfile))
		v1Router.Delete("/users/profile", authHandlers.MiddlewareAuth(usersHandlers.DeleteProfile))

		rolesHandlers := controllers.NewRolesHandlers(configuration.ApiCfg.DB)

		v1Router.Get("/roles", authHandlers.MiddlewareAuth(rolesHandlers.GetAll))
		v1Router.Post("/roles", authHandlers.MiddlewareAuth(rolesHandlers.Create))
		v1Router.Get("/roles/{id}", authHandlers.MiddlewareAuth(rolesHandlers.Get))
		v1Router.Put("/roles/{id}", authHandlers.MiddlewareAuth(rolesHandlers.Update))
		v1Router.Put("/roles/{id}", authHandlers.MiddlewareAuth(rolesHandlers.Delete))

		genresHandlers := controllers.NewGenresHandlers(configuration.ApiCfg.DB)

		v1Router.Get("/genres", authHandlers.MiddlewareAuth(genresHandlers.GetAll))
		v1Router.Post("/genres", authHandlers.MiddlewareAuth(genresHandlers.Create))
		v1Router.Get("/genres/{id}", authHandlers.MiddlewareAuth(genresHandlers.Get))
		v1Router.Put("/genres/{id}", authHandlers.MiddlewareAuth(genresHandlers.Update))
		v1Router.Put("/genres/{id}", authHandlers.MiddlewareAuth(genresHandlers.Delete))

		ageCategoriesHandlers := controllers.NewAgeCategoriesHandlers(configuration.ApiCfg.DB)

		v1Router.Get("/age-categories", authHandlers.MiddlewareAuth(ageCategoriesHandlers.GetAll))
		v1Router.Post("/age-categories", authHandlers.MiddlewareAuth(ageCategoriesHandlers.Create))
		v1Router.Get("/age-categories/{id}", authHandlers.MiddlewareAuth(ageCategoriesHandlers.Get))
		v1Router.Put("/age-categories/{id}", authHandlers.MiddlewareAuth(ageCategoriesHandlers.Update))
		v1Router.Put("/age-categories/{id}", authHandlers.MiddlewareAuth(ageCategoriesHandlers.Delete))

		typesHandlers := controllers.NewTypesHandlers(configuration.ApiCfg.DB)

		v1Router.Get("/types", authHandlers.MiddlewareAuth(typesHandlers.GetAll))
		v1Router.Post("/types", authHandlers.MiddlewareAuth(typesHandlers.Create))
		v1Router.Get("/types/{id}", authHandlers.MiddlewareAuth(typesHandlers.Get))
		v1Router.Put("/types/{id}", authHandlers.MiddlewareAuth(typesHandlers.Update))
		v1Router.Put("/types/{id}", authHandlers.MiddlewareAuth(typesHandlers.Delete))

		imagesHandlers := controllers.NewImagesHandlers(configuration.ApiCfg.DB, configuration.ApiCfg.Dir)

		v1Router.Post("/projects/images", authHandlers.MiddlewareAuth(imagesHandlers.Upload))
		v1Router.Get("/projects/images/{id}", authHandlers.MiddlewareAuth(imagesHandlers.Get))
		v1Router.Get("/projects/images/show/{id}", authHandlers.MiddlewareAuth(imagesHandlers.Display))
		v1Router.Delete("/projects/images/{id}", authHandlers.MiddlewareAuth(imagesHandlers.Delete))

		videosHandlers := controllers.NewVideosHandlers(configuration.ApiCfg.DB, configuration.ApiCfg.Dir)

		v1Router.Post("/projects/videos", authHandlers.MiddlewareAuth(videosHandlers.Upload))
		v1Router.Get("/projects/videos/{id}", authHandlers.MiddlewareAuth(videosHandlers.Get))
		v1Router.Delete("/projects/videos/{id}", authHandlers.MiddlewareAuth(videosHandlers.Delete))
		v1Router.Get("/projects/videos/play/{id}", authHandlers.MiddlewareAuth(videosHandlers.Play))

		projectsRepository := repositories.NewProjectsRepository(configuration.ApiCfg.Conn)
		projectsHandlers := controllers.NewProjecsHandlers(projectsRepository, configuration.ApiCfg.Dir)

		v1Router.Get("/projects", authHandlers.MiddlewareAuth(projectsHandlers.GetAll))
		v1Router.Post("/projects", authHandlers.MiddlewareAuth(projectsHandlers.Create))
		v1Router.Put("/projects/{id}", authHandlers.MiddlewareAuth(projectsHandlers.Update))

		v1Router.Post("/projects/{id}/cover", authHandlers.MiddlewareAuth(projectsHandlers.UploadCover))
		v1Router.Patch("/projects/{id}/cover", authHandlers.MiddlewareAuth(projectsHandlers.SetCover))
	}

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 10,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
