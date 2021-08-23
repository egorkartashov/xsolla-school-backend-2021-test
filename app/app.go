package app

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/controllers"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type App struct {
	Router                    *mux.Router
	productsController        *controllers.ProductsController
	productsGraphqlController *controllers.ProductsGraphqlController
	logger                    *log.Entry
}

func New(db *gorm.DB, logger *log.Entry) (*App, error) {
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	var productsRepo repos.ProductsRepoInterface = repos.NewProductsRepo(db)
	messageBrokerService := services.NewMessageBrokerService()
	productsService := services.NewProductsService(productsRepo, messageBrokerService)

	a := &App{
		Router:                    mux.NewRouter(),
		productsController:        controllers.NewProductsController(productsService, logger),
		productsGraphqlController: controllers.NewProductsGraphqlController(productsService),
		logger:                    logger,
	}

	a.registerHandlers()

	return a, nil
}

func (a *App) registerHandlers() {
	a.Router.Path("/api/products").Queries("query", "{.*}").HandlerFunc(a.productsGraphqlController.HandleQuery)
	a.Router.HandleFunc("/api/ping", controllers.GetPing)

	productsRouter := a.Router.PathPrefix("/api/products").Subrouter()
	productsRouter.Use(a.logApiRequestMW)

	productsRouter.HandleFunc("/types", a.productsController.GetAllTypes).Methods("GET")
	productsRouter.HandleFunc("", a.productsController.GetProducts).Methods("GET")
	productsRouter.HandleFunc("", a.productsController.PostProduct).Methods("POST")
	productsRouter.HandleFunc("/sku={sku}", a.productsController.GetProductBySku).Methods("GET")
	productsRouter.HandleFunc("/sku={sku}", a.productsController.PutProductBySku).Methods("PUT")
	productsRouter.HandleFunc("/sku={sku}", a.productsController.DeleteProductBySku).Methods("DELETE")
	productsRouter.HandleFunc("/{id}", a.productsController.GetProduct).Methods("GET")
	productsRouter.HandleFunc("/{id}", a.productsController.PutProduct).Methods("PUT")
	productsRouter.HandleFunc("/{id}", a.productsController.DeleteProduct).Methods("DELETE")
}

func (a *App) logApiRequestMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathTemplate, _ := mux.CurrentRoute(r).GetPathTemplate()
		httpRequestInfo := log.Fields{
			"method":       r.Method,
			"pathTemplate": pathTemplate,
			"timestampUtc": time.Now().UTC(),
		}

		a.logger.WithFields(log.Fields{"httpRequestInfo": httpRequestInfo}).Info("REST API request")

		next.ServeHTTP(w, r)
	})
}
