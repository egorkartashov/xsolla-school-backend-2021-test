package controllers

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/dto"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/filters"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/graphql-api"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type ProductsGraphqlController struct {
	queryResolver ProductsGraphqlQueryResolver
	schema        graphql.Schema
}

type ProductsGraphqlQueryResolver struct {
	productsService *services.ProductsService
}

func NewProductsGraphqlController(productsService *services.ProductsService) *ProductsGraphqlController {
	queryResolver := ProductsGraphqlQueryResolver{
		productsService: productsService,
	}

	return &ProductsGraphqlController{
		queryResolver: queryResolver,
		schema:        CreateProductGraphqlSchema(queryResolver),
	}
}

func CreateProductGraphqlSchema(queryResolver ProductsGraphqlQueryResolver) graphql.Schema {
	var productType = graphql_api.GetProductType()
	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queryResolver.getProductById,
			},
			"productBySku": &graphql.Field{
				Type:        productType,
				Description: "Get product by SKU",
				Args: graphql.FieldConfigArgument{
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queryResolver.getProductBySku,
			},
			"productsList": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get products list",
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: queryResolver.getProductsList,
			},
		},
	})

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        productType,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"priceInCents": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: queryResolver.createProduct,
			},
			"update": &graphql.Field{
				Type:        productType,
				Description: "Update product by its ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"priceInCents": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: queryResolver.updateProductById,
			},
			"updateBySku": &graphql.Field{
				Type:        productType,
				Description: "Update product by its SKU",
				Args: graphql.FieldConfigArgument{
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"priceInCents": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: queryResolver.updateProductBySku,
			},
			"delete": &graphql.Field{
				Type:        graphql.String,
				Description: "Delete product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queryResolver.deleteProductById,
			},
			"deleteBySku": {
				Type:        graphql.String,
				Description: "Delete product by SKU",
				Args: graphql.FieldConfigArgument{
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queryResolver.deleteProductBySku,
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
	result := graphql.Do(graphql.Params{
		Schema:        controller.schema,
		RequestString: queryString,
	})

	if len(result.Errors) > 0 {
		log.Printf("errors: %v", result.Errors)
		responseBody := map[string]interface{}{
			"data":   "Error during /products GraphQL query",
			"errors": result.Errors,
		}
		utils.RespondJson(w, http.StatusInternalServerError, responseBody)
		return
	}

	utils.RespondJson(w, http.StatusOK, result)
}

func (queryResolver ProductsGraphqlQueryResolver) getProductById(p graphql.ResolveParams) (interface{}, error) {
	idString, _ := p.Args["id"].(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	product, requestResult := queryResolver.productsService.GetProduct(id)
	if requestResult.Status == services.Success {
		return product, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) getProductBySku(p graphql.ResolveParams) (interface{}, error) {
	sku, _ := p.Args["sku"].(string)
	product, requestResult := queryResolver.productsService.GetProductBySku(sku)
	if requestResult.Status == services.Success {
		return product, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) getProductsList(params graphql.ResolveParams) (interface{}, error) {
	var offset, limit int
	var isArgSpecified bool

	if offset, isArgSpecified = params.Args["offset"].(int); !isArgSpecified {
		offset = 0
	}
	if limit, isArgSpecified = params.Args["limit"].(int); !isArgSpecified {
		limit = 50
	}

	log.Printf("%v, %v", offset, limit)

	var noFilters = make([]filters.FilterPair, 0)
	products, requestResult := queryResolver.productsService.GetProducts(noFilters, offset, limit)
	if requestResult.Status == services.Success {
		return products, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) createProduct(params graphql.ResolveParams) (interface{}, error) {
	newProduct := dto.ProductDto{
		Sku:          params.Args["sku"].(string),
		Name:         params.Args["name"].(string),
		Type:         params.Args["type"].(string),
		PriceInCents: int32(params.Args["priceInCents"].(int)),
	}

	productDto, requestResult := queryResolver.productsService.CreateProduct(&newProduct)
	if requestResult.Status == services.Success {
		return productDto, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) updateProductById(params graphql.ResolveParams) (interface{}, error) {
	idString, _ := params.Args["id"].(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		return "ok", err
	}

	productToUpdate, _ := queryResolver.productsService.GetProduct(id)

	sku, skuChanged := params.Args["sku"].(string)
	if skuChanged {
		productToUpdate.Sku = sku
	}

	name, nameChanged := params.Args["name"].(string)
	if nameChanged {
		productToUpdate.Name = name
	}

	productType, typeChanged := params.Args["type"].(string)
	if typeChanged {
		productToUpdate.Type = productType
	}

	priceInCents, priceChanged := params.Args["priceInCents"].(int32)
	if priceChanged {
		productToUpdate.PriceInCents = priceInCents
	}

	updatedProduct, requestResult := queryResolver.productsService.UpdateProduct(productToUpdate)
	if requestResult.Status == services.Success {
		return updatedProduct, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) updateProductBySku(params graphql.ResolveParams) (interface{}, error) {
	sku, _ := params.Args["sku"].(string)
	productToUpdate, _ := queryResolver.productsService.GetProductBySku(sku)

	name, nameChanged := params.Args["name"].(string)
	if nameChanged {
		productToUpdate.Name = name
	}

	productType, typeChanged := params.Args["type"].(string)
	if typeChanged {
		productToUpdate.Type = productType
	}

	priceInCents, priceChanged := params.Args["priceInCents"].(int32)
	if priceChanged {
		productToUpdate.PriceInCents = priceInCents
	}

	updatedProduct, requestResult := queryResolver.productsService.UpdateProduct(productToUpdate)
	if requestResult.Status == services.Success {
		return updatedProduct, nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) deleteProductBySku(params graphql.ResolveParams) (interface{}, error) {
	sku, _ := params.Args["sku"].(string)

	requestResult := queryResolver.productsService.DeleteProductBySku(sku)
	if requestResult.Status == services.Success {
		return fmt.Sprintf("Deleted product with SKU=%v", sku), nil
	} else {
		return nil, requestResult.Error
	}
}

func (queryResolver ProductsGraphqlQueryResolver) deleteProductById(params graphql.ResolveParams) (interface{}, error) {
	idString, _ := params.Args["id"].(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		return "ok", err
	}

	requestResult := queryResolver.productsService.DeleteProduct(id)
	if requestResult.Status == services.Success {
		return fmt.Sprintf("Deleted product with ID=%v", id), nil
	} else {
		return nil, requestResult.Error
	}
}
