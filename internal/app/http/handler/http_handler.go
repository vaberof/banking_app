package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/vaberof/MockBankingApplication/docs"
)

type HttpHandler struct {
	userService     UserService
	accountService  AccountService
	transferService TransferService
	depositService  DepositService
	authService     AuthorizationService
}

func NewHttpHandler(
	userService UserService,
	accountService AccountService,
	transferService TransferService,
	depositService DepositService,
	authService AuthorizationService) *HttpHandler {

	return &HttpHandler{
		userService:     userService,
		accountService:  accountService,
		transferService: transferService,
		depositService:  depositService,
		authService:     authService,
	}
}

func (h *HttpHandler) InitRoutes(config *fiber.Config) *fiber.App {
	app := fiber.New(*config)

	h.configureCors(app)

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	{
		api.Post("/register", h.register)
		api.Post("/login", h.login)
		api.Post("/logout", h.logout)

		api.Post("/account", h.createAccount)
		api.Delete("/account", h.deleteAccount)
		api.Get("/accounts", h.getAccounts)

		api.Post("/transfer", h.makeTransfer)
		api.Get("/transfers", h.getTransfers)

		api.Get("/deposits", h.getDeposits)
	}

	return app
}

func (h *HttpHandler) configureCors(app *fiber.App) {
	corsConfig := cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}

	app.Use(cors.New(corsConfig))
}
