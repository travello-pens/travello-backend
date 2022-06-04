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

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlUserRepository(db)
	service := u.NewServiceUser(repo, conf)
	hand := h.EchoUserController{Service: service}

	apiUser := e.Group("/user",
		middleware.Logger(),
		middleware.CORS(),
	)

	// apiUser.GET("", hand.GetUsersController, middleware.JWTWithConfig(
	// 	middleware.JWTConfig{
	// 		SigningKey: []byte(conf.JWT_KEY),
	// 		ErrorHandlerWithContext: func(err error, c echo.Context) error {
	// 			return c.JSONPretty(404, map[string]interface{}{
	// 				"messages": "token tidak valid",
	// 			}, " ")
	// 		},
	// 		SuccessHandler: func(c echo.Context) {},
	// 	},
	// ))
	apiUser.GET("", hand.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.GET("/:id", hand.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", hand.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/:id", hand.DeleteUsercontroller, middleware.JWT([]byte(conf.JWT_KEY)))

}
