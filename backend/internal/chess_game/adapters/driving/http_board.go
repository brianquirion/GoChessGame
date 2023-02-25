package driving

import (
	"github.com/gofiber/fiber/v2"
	"goChessGame/internal/chess_game/core/command_handlers"
	"strconv"
)

func GetBoard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(fiber.Map{
		"board": command_handlers.GetCurrentBoard(id),
	})
}

func PostBoardMove(c *fiber.Ctx) error {
	payload := struct {
		PositionFrom string `json:"positionFrom"`
		PositionTo   string `json:"positionTo"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	return c.JSON("")
}
