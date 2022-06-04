package route

import (
	"travel-agent-backend/config"
	d "travel-agent-backend/database"
	h "travel-agent-backend/internal/handler"
	r "travel-agent-backend/internal/repository"
	u "travel-agent-backend/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAuthGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlAuthRepository(db)
	service := u.NewServiceAuth(repo, conf)
	hand := h.EchoAuthController{Service: service}

	e.POST("/register", hand.RegisterController, middleware.Logger(), middleware.CORS())
	e.POST("/login", hand.LoginController, middleware.Logger(), middleware.CORS())
}
