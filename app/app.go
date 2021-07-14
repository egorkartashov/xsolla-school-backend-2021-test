package app

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/controllers"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router             *mux.Router
	productsController *controllers.ProductsController
}

func New(db *gorm.DB) (*App, error) {
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	productsRepo := repos.NewProductsRepo(db)
	productsService := services.NewProductsService(productsRepo)

	a := &App{
		Router:             mux.NewRouter(),
		productsController: controllers.NewProductsController(productsService),
	}

	a.registerHandlers()

	return a, nil
}

func (a *App) registerHandlers() {
	a.Router.HandleFunc("/api/ping", controllers.GetPing)
	a.Router.HandleFunc("/api/products/{id}", a.productsController.GetProduct).Methods("GET")
	a.Router.HandleFunc("/api/products", a.productsController.PostProduct).Methods("POST")
	a.Router.HandleFunc("/api/products/{id}", a.productsController.PutProduct).Methods("PUT")
	a.Router.HandleFunc("/api/products/{id}", a.productsController.DeleteProductById).Methods("DELETE")
	a.Router.HandleFunc("/api/products", a.productsController.DeleteProductBySku).Methods("DELETE")
	a.Router.HandleFunc("/api/products/types", a.productsController.GetAllTypes).Methods("GET")
}
