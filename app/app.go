package app

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/auth"
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

	productsRouter := a.Router.PathPrefix("/api/products").Subrouter()
	productsRouter.Use(auth.JwtAuthMW)
	modifyProductsRouter := a.Router.PathPrefix("/api/products").Subrouter()
	modifyProductsRouter.Use(auth.VendorAuthMW)

	modifyProductsRouter.HandleFunc("", a.productsController.PostProduct).Methods("POST")
	productsRouter.HandleFunc("/types", a.productsController.GetAllTypes).Methods("GET")
	productsRouter.HandleFunc("", a.productsController.GetProducts).Methods("GET")
	productsRouter.HandleFunc("/sku={sku}", a.productsController.GetProductBySku).Methods("GET")
	modifyProductsRouter.HandleFunc("/sku={sku}", a.productsController.PutProductBySku).Methods("PUT")
	modifyProductsRouter.HandleFunc("/sku={sku}", a.productsController.DeleteProductBySku).Methods("DELETE")
	productsRouter.HandleFunc("/{id}", a.productsController.GetProduct).Methods("GET")
	modifyProductsRouter.HandleFunc("/{id}", a.productsController.PutProduct).Methods("PUT")
	modifyProductsRouter.HandleFunc("/{id}", a.productsController.DeleteProduct).Methods("DELETE")
}
