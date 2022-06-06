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

func RegisterTravelAgentSocialMediaGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlTravelAgentSocialMediaRepository(db)
	service := u.NewServiceTravelAgentSocialMedia(repo, conf)
	hand := h.EchoTravelAgentSocialMediaController{Service: service}

	apiTravelAgentSocialMedia := e.Group("/travel-agent-social-media",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiTravelAgentSocialMedia.POST("", hand.CreateTravelAgentSocialMediaController)
	apiTravelAgentSocialMedia.GET("", hand.GetTravelAgentSocialMediasController)
	apiTravelAgentSocialMedia.GET("/:id", hand.GetTravelAgentSocialMediaController)
	apiTravelAgentSocialMedia.PUT("/:id", hand.UpdateTravelAgentSocialMediaController)
	apiTravelAgentSocialMedia.DELETE("/:id", hand.DeleteTravelAgentSocialMediaController)
}
