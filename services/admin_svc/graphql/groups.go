package graphql

import (
	"github.com/graphql-go/graphql"
)

var fields = graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "world", nil
	},
}
