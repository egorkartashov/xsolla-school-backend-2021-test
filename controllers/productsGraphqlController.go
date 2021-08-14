package controllers

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/graphql-api"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"net/http"
)

type ProductsGraphqlController struct {
	productsService *services.ProductsService
}

func NewProductsGraphqlController(productsService *services.ProductsService) *ProductsGraphqlController {
	return &ProductsGraphqlController{
		productsService: productsService,
	}
}

func CreateProductGraphqlSchema(productsService *services.ProductsService) graphql.Schema {
	var productType = graphql_api.GetProductType()
	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"sku": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idString, foundId := p.Args["id"].(string)
					if foundId {
						id, parsingError := uuid.Parse(idString)
						if parsingError == nil {
							product, requestResult := productsService.GetProduct(id)
							if requestResult.Status == services.Success {
								return product, nil
							} else {
								return nil, requestResult.Error
							}
						}

						return nil, nil // TODO return error message
					}

					sku, foundSku := p.Args["sku"].(string)
					if !foundId && !foundSku {
						product, requestResult := productsService.GetProductBySku(sku)
						if requestResult.Status == services.Success {
							return product, nil
						} else {
							return nil, requestResult.Error
						}
					}

					return nil, nil // TODO return error message
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get products list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					products, requestResult := productsService.GetProducts(0, 50, nil)
					if requestResult.Status == services.Success {
						return products, nil
					} else {
						return nil, requestResult.Error
					}
				},
			},
		},
	})

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				//TODO
			},
			"update": &graphql.Field{
				//TODO
			},
			"delete": &graphql.Field{
				//TODO
			},
		},
	})

	var productGraphqlSchema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)

	return productGraphqlSchema
}

func (controller *ProductsGraphqlController) HandleQuery(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query().Get("query")
	schema := CreateProductGraphqlSchema(controller.productsService)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: queryString,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
		utils.RespondErrorJson(w, http.StatusInternalServerError, "Error during /products GraphQL query")
		return
	}

	utils.RespondJson(w, http.StatusOK, result)
}
