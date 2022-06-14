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

func RegisterProductGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlProductRepository(db)
	service := u.NewServiceProduct(repo, conf)
	hand := h.EchoProductController{Service: service}

	apiProduct := e.Group("/product",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiProduct.POST("", hand.CreateProductController)
	apiProduct.GET("", hand.GetProductsController)
	apiProduct.GET("/:id", hand.GetProductController)
	apiProduct.PUT("/:id", hand.UpdateProductController)
	apiProduct.DELETE("/:id", hand.DeleteProductController)
	apiProduct.GET("/agent/:agent", hand.GetProductsByAgentController)
	apiProduct.GET("/location/:location", hand.GetProductsByLocationController)
}
