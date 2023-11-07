package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Config struct {
	HostPort  uint
	Router    *chi.Mux
	StripeKey string
	Domain    string
}

type App struct {
	*Config
	Server *http.Server
}

func New(config *Config) *App {

	if config == nil {
		config = &Config{}
	}

	if config.Router == nil {
		config.Router = chi.NewRouter()
	}

	if config.HostPort == 0 {
		config.HostPort = 8080
	}

	if config.Domain == "" {
		config.Domain = "http://localhost:8080"
	}

	a := &App{
		Config: config,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.HostPort),
			Handler: config.Router,
		},
	}

	a.initializeRoutes()

	return a
}

func (a *App) Run() {
	fmt.Printf("Listening on port %s\n", a.Server.Addr)
	if err := a.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (a *App) initializeRoutes() {
	a.Router.Get("/", a.handleIndex)
	a.Router.Post("/create-checkout-session", a.handleCreateCheckoutSession)
	a.Router.Mount("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
}

func (a *App) PrintRoutes() {
	// 👇 the walking function 🚶‍♂️
	chi.Walk(a.Router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
}
