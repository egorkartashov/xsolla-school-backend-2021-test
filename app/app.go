package app

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/controllers"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router             *mux.Router
	productsController *controllers.ProductsController
}

func New(db *gorm.DB) *App {
	productsRepo := repos.NewProductsRepo(db)
	productsService := services.NewProductsService(productsRepo)

	a := &App{
		Router:             mux.NewRouter(),
		productsController: controllers.NewProductsController(productsService),
	}

	a.registerHandlers()

	return a
}

func (a *App) registerHandlers() {
	a.Router.HandleFunc("/api/ping", controllers.GetPing)
	a.Router.HandleFunc("/api/products/{id}", a.productsController.GetProduct)
}
