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

func RegisterOrderGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlOrderRepository(db)
	service := u.NewServiceOrder(repo, conf)
	hand := h.EchoOrderController{Service: service}

	apiOrder := e.Group("/order",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiOrder.POST("", hand.CreateOrderController)
	apiOrder.GET("", hand.GetOrdersController)
	apiOrder.GET("/:id", hand.GetOrderController)
	apiOrder.PUT("/:id", hand.UpdateOrderController)
	apiOrder.DELETE("/:id", hand.DeleteOrderController)
}
