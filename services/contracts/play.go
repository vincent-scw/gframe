package contracts

// Shape enum
type Shape int

const (
	// Rock shape
	Rock Shape = 1
	// Paper shape
	Paper Shape = 2
	// Scissors shape
	Scissors Shape = 3
)

// Play a round
type Play struct {
	Player User  `json:"player"`
	Shape  Shape `json:"shape"`
}

// Game one game
type Game struct {
	Play  Play      `json:"play"`
	Group GroupInfo `json:"group"`
}
