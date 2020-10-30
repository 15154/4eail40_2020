package piece

import (
	"github.com/chess/model/coord"
	"github.com/chess/model/player"
)

type Rook struct {
}

func (r Rook) String() string {
	return ""
}
func (r Rook) Color() player.Color {
	return player.White
}
func (r Rook) Moves(isCapture bool) map[coord.ChessCoordinates]bool {
	return nil
}
