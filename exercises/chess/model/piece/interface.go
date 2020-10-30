// Package piece will handle game pieces.
package piece

import (
	"fmt"

	"github.com/chess/model/coord"
	"github.com/chess/model/player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns the appartenance.
	Color() player.Color
	// Moves returns a set of valid move.
	Moves(isCapture bool) map[coord.ChessCoordinates]bool
}
