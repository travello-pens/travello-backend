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

func RegisterTravelAgentGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlTravelAgentRepository(db)
	service := u.NewServiceTravelAgent(repo, conf)
	hand := h.EchoTravelAgentController{Service: service}

	apiTravelAgent := e.Group("/agent",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiTravelAgent.POST("", hand.CreateTravelAgentController)
	apiTravelAgent.GET("", hand.GetTravelAgentsController)
	apiTravelAgent.GET("/:id", hand.GetTravelAgentController)
	apiTravelAgent.PUT("/:id", hand.UpdateTravelAgentController)
	apiTravelAgent.DELETE("/:id", hand.DeleteTravelAgentController)
	apiTravelAgent.GET("/product", hand.GetProductsByAgentTravelController)
	apiTravelAgent.GET("/sum-product/:agent", hand.GetSumProductByAgentController)
}
