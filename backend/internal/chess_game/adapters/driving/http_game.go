package driving

import (
	"github.com/gofiber/fiber/v2"
	"goChessGame/internal/chess_game/core/domain/constants"
	"goChessGame/internal/chess_game/core/domain/enums"
)

func GetBoard(c *fiber.Ctx) error {
	newBoard := [constants.BOARD_HEIGHT][constants.BOARD_WIDTH]enums.Piece{
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

	return c.JSON(fiber.Map{
		"board": newBoard,
	})
}

func PostBoardMove(c *fiber.Ctx) error {
	return c.JSON("")
}
