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

func RegisterProductPhotoGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlProductPhotoRepository(db)
	service := u.NewServiceProductPhoto(repo, conf)
	hand := h.EchoProductPhotoController{Service: service}

	apiProductPhoto := e.Group("/product-photo",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiProductPhoto.POST("", hand.CreateProductPhotoController)
	apiProductPhoto.GET("", hand.GetProductPhotosController)
	apiProductPhoto.GET("/:id", hand.GetProductPhotoController)
	apiProductPhoto.PUT("/:id", hand.UpdateProductPhotoController)
	apiProductPhoto.DELETE("/:id", hand.DeleteProductPhotoController)
	apiProductPhoto.GET("/product/:product", hand.GetProductPhotoByProductController)
}
