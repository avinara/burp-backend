package app

import (
	"fmt"
	"net/http"

	"github.com/burp-backend/config"
	"github.com/burp-backend/controllers"
	"github.com/burp-backend/database"
	"github.com/burp-backend/errors"
	"github.com/burp-backend/routes"
	"github.com/burp-backend/services"
)

type App struct {
	Config config.Config
	Router routes.RouterInterface
}

func NewApp() *App {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		err2 := errors.LoadingConfigurationFileError()
		panic(err2)
	}

	return &App{
		Config: *config,
	}
}

func (app *App) Init() {
	db, err := database.InitDB(&app.Config.DatabaseConfig)
	if err != nil {
		err2 := errors.DatabaseInitError()
		panic(err2)
	}

	cookService := services.NewCookService(db)
	userService := services.NewUserService(db)
	cookController := controllers.NewCookController(cookService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(app.Config)
	router := routes.NewRouter()
	app.Router = router
	router.InitRoutes(cookController, userController, authController)
}

func (app *App) Run() {
	fmt.Println("Starting server on port 8080")
	fmt.Println("############## Server Started ##############")
	http.ListenAndServe(":8080", app.Router.GetMux())

}
