package routes

import (
	"net/http"

	"github.com/burp-backend/config"
	"github.com/burp-backend/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RouterInterface interface {
	InitRoutes(cookController controllers.CookControllerAPI, userController controllers.UserControllerAPI, authController controllers.AuthControllerAPI)
	GetMux() *chi.Mux
}

type router struct {
	mux *chi.Mux
}

func NewRouter(config config.Config) RouterInterface {
	mux := chi.NewRouter()
	mux.Use(cors.New(cors.Options{
		AllowedOrigins: config.FrontendConfig.FrontendURL,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Accept-Encoding", "Accept-Language", "Connection", "Cookie", "Host",
			"Origin", "Referer", "User-Agent", "Content-Type"},
		AllowCredentials: true,
	}).Handler)

	return &router{
		mux: mux,
	}
}

func (h *router) GetMux() *chi.Mux {
	return h.mux
}

func (h *router) InitRoutes(cookController controllers.CookControllerAPI, userController controllers.UserControllerAPI, authController controllers.AuthControllerAPI) {

	h.mux.Group(func(r chi.Router) {

		// cook crud routes
		r.Get("/cooks", cookController.GetAllCooks)
		r.Post("/cook", cookController.CreateCook)
		r.Post("/cook/update", cookController.UpdateCook)
		r.Delete("/cook/delete", cookController.DeleteCook)
		r.Get("/cook", cookController.GetCookByEmail)

		// user crud routes
		r.Post("/user", userController.CreateUser)
		r.Post("/user/update", userController.UpdateUser)
		r.Delete("/user/delete", userController.DeleteUser)
		r.Get("/user/{name}", userController.GetUserByName)

		r.Get("/google_login", authController.GoogleLogin)
		r.Get("/google_callback", authController.GoogleCallback)
	})

	h.mux.With(RemoveContextTypeJSON).Get("/swagger/*", httpSwagger.WrapHandler)
}

func RemoveContextTypeJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
