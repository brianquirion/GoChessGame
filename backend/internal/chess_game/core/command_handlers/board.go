package command_handlers

import (
	"goChessGame/internal/chess_game/core/domain/enums"
	"goChessGame/internal/chess_game/core/domain/models"
)

func GetNewBoard() models.Board {
	newBoard := models.Board{
		{enums.BLACK_ROOK, enums.BLACK_KNIGHT, enums.BLACK_BISHOP, enums.BLACK_QUEEN, enums.BLACK_KING,
			enums.BLACK_BISHOP, enums.BLACK_KNIGHT, enums.BLACK_ROOK},
		{enums.BLACK_PAWN, enums.BLACK_PAWN, enums.BLACK_PAWN, enums.BLACK_PAWN, enums.BLACK_PAWN, enums.BLACK_PAWN,
			enums.BLACK_PAWN, enums.BLACK_PAWN},
		{},
		{},
		{},
		{},
		{enums.WHITE_PAWN, enums.WHITE_PAWN, enums.WHITE_PAWN, enums.WHITE_PAWN, enums.WHITE_PAWN, enums.WHITE_PAWN,
			enums.WHITE_PAWN, enums.WHITE_PAWN},
		{enums.WHITE_ROOK, enums.WHITE_KNIGHT, enums.WHITE_BISHOP, enums.WHITE_QUEEN, enums.WHITE_KING,
			enums.WHITE_BISHOP, enums.WHITE_KNIGHT, enums.WHITE_ROOK},
	}
	return newBoard
}

func GetCurrentBoard(id int) models.Board {
	return GetNewBoard()
}
