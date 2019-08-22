package graphql

import (
	"testing"

	"github.com/graphql-go/graphql"
)

func TestInjectPlayers(t *testing.T) {
	injectPlayersField.Resolve = dummyInjectPlayersHandler

	mutation := `
	mutation simulator($amount: Int!) {
		injectPlayers(amount: $amount) {
			amount
		}
	}
	`
	variables := make(map[string]interface{})
	variables["amount"] = 2

	params := graphql.Params{
		Schema:         Schema,
		RequestString:  mutation,
		VariableValues: variables,
		OperationName:  "simulator",
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
}

func dummyInjectPlayersHandler(params graphql.ResolveParams) (interface{}, error) {
	return InjectPlayers{Amount: 1}, nil
}
