package graphql

import (
	"testing"

	"github.com/graphql-go/graphql"
)

func TestInjectPlayers(t *testing.T) {
	injectPlayersField.Resolve = dummyInjectPlayersHandler

	mutation := `
	mutation Simulator {
		injectPlayers(amount: 4) {
			amount
		}
	}
	`

	params := graphql.Params{Schema: Schema, RequestString: mutation}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
}

func dummyInjectPlayersHandler(params graphql.ResolveParams) (interface{}, error) {
	return InjectPlayers{Amount: 1}, nil
}
