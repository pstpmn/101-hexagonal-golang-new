package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	ports2 "learn-oauth2/internal/core/ports"
)

type server struct {
	membersHandler ports2.IMembersHandler
	middlewares    ports2.IMiddlewares
	env            map[string]interface{}
}

func NewServer(membersHandler ports2.IMembersHandler, middlewares ports2.IMiddlewares, env map[string]interface{}) ports2.IServer {
	return &server{membersHandler: membersHandler, middlewares: middlewares, env: env}
}

func (s server) Initialize() {
	app := fiber.New()

	// init middleware
	s.middleware(app)

	// init routes
	s.routes(app)

	// init app
	if err := app.Listen(fmt.Sprintf(":%d", s.env["PORT"])); err != nil {
		panic(err)
	}
}

func (s server) routes(app *fiber.App) {
	app.Post("registration", s.membersHandler.Registration)
	app.Post("authentication", s.membersHandler.Authentication)

	// api authorize
	authz := app.Group("/authorization")
	authz.Get("/", s.membersHandler.Authorization)
	authz.Get("/facebook/:accessToken", s.membersHandler.AuthorizationForFacebook)
	authz.Get("/google", s.membersHandler.AuthorizationForGoogle)

	// api for pass authorize
	//auth := app.Group("/auth", s.middlewares.Authorize)
	//auth.Get("/", s.membersHandler.HelloWorld)
}

func (s server) middleware(app *fiber.App) {
}
