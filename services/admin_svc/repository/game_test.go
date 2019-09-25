package repository

import (
	"testing"
	"time"
	//c "github.com/vincent-scw/gframe/contracts"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewGame(t *testing.T) {
	model := NewGame("test game", "vincent", time.Time{})
	if model.Name != "test game" || model.CreatedBy != "vincent" || !model.RegisterTime.IsZero() {
		t.Error("NewGame model error")
	}
}

// func TestCreateGame(t *testing.T) {
// 	repo := NewGameRepository()
// 	one, err := repo.GetOne("test321")
// 	if err != nil {
// 		if one != nil {

// 		}
// 	}
// }
