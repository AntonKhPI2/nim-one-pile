package game

type Variant string

const (
	Normal Variant = "normal"
	Misere Variant = "misere"
)

type Config struct {
	Variant Variant `json:"variant"`
	N       int     `json:"n"`
	K       int     `json:"k"`
}

type State struct {
	ID         string  `json:"id"`
	Variant    Variant `json:"variant"`
	K          int     `json:"k"`
	Remaining  int     `json:"remaining"`
	PlayerTurn string  `json:"player_turn"`
	Winner     string  `json:"winner"`
}

type MoveRequest struct {
	ID   string `json:"id"`
	Take int    `json:"take"`
}
