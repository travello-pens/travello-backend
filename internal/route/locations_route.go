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

func RegisterLocationGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlLocationRepository(db)
	service := u.NewServiceLocation(repo, conf)
	hand := h.EchoLocationController{Service: service}

	apiLocation := e.Group("/location",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiLocation.POST("", hand.CreateLocationController)
	apiLocation.GET("", hand.GetLocationsController)
	apiLocation.GET("/:id", hand.GetLocationController)
	apiLocation.PUT("/:id", hand.UpdateLocationController)
	apiLocation.DELETE("/:id", hand.DeleteLocationController)
}
