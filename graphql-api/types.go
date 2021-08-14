package graphql_api

import (
	"github.com/graphql-go/graphql"
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"priceInCents": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func GetProductType() *graphql.Object {
	return productType
}
