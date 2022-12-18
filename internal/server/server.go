package server

import (
	"fmt"
	"lean-oauth/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	membersHandler ports.IMembersHandler
	env            map[string]interface{}
}

func NewServer(membersHandler ports.IMembersHandler, env map[string]interface{}) ports.IServer {
	return &server{membersHandler: membersHandler, env: env}
}

func (s server) Initialize() {
	app := fiber.New()

	//init middleware
	s.middleware(app)

	// init routes
	s.routes(app)

	// init app
	app.Listen(fmt.Sprintf(":%d", s.env["PORT"]))
}

func (s server) routes(app *fiber.App) {
	app.Get("/", s.membersHandler.HelloWorld)
	app.Post("registration", s.membersHandler.Registration)
}

func (s server) middleware(app *fiber.App) {
}
