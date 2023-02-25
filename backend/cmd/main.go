package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goChessGame/internal/chess_game/adapters/driving"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	fmt.Print("Starting up...")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	v1.Get("/ping", driving.Ping)
	v1.Get("/board/:id", driving.GetBoard)
	v1.Post("/board/:id/move", driving.PostBoardMove)

	app.Static("/", "./static/public")

	app.Use(driving.NotFound)

	log.Fatal(app.Listen(*port))
}
