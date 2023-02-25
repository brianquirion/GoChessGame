package models

import (
	"goChessGame/internal/chess_game/core/domain/constants"
	"goChessGame/internal/chess_game/core/domain/enums"
)

type Board = [constants.BOARD_HEIGHT][constants.BOARD_WIDTH]enums.Piece
