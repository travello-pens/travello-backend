package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoOrderController struct {
	Service adapter.AdapterOrderService
}

// CreateOrderController godoc
// @Summary      Create Order
// @Description  User can create order
// @Tags         Order
// @accept       json
// @Produce      json
// @Router       /order [post]
// @param        data  body      model.Order  true  "required"
// @Success      201   {object}  model.Order
// @Failure      500   {object}  model.Order
// @Security     JWT
func (ce *EchoOrderController) CreateOrderController(c echo.Context) error {
	order := model.Order{}
	c.Bind(&order)

	err := ce.Service.CreateOrderService(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"order":    order,
	})
}

// GetOrdersController godoc
// @Summary      Get All Orders Information
// @Description  User can get all orders information
// @Tags         Order
// @accept       json
// @Produce      json
// @Router       /order [get]
// @Success      200   {object}  model.Order
// @Security     JWT
func (ce *EchoOrderController) GetOrdersController(c echo.Context) error {
	orders := ce.Service.GetAllOrdersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"orders":   orders,
	}, " ")
}

// GetOrderController godoc
// @Summary      Get Order Information by Id
// @Description  User can get order information by id
// @Tags         Order
// @accept       json
// @Produce      json
// @Router       /order/{id} [get]
// @param        id    path      int          true  "id"
// @Success      200  {object}  model.Order
// @Failure      404  {object}  model.Order
// @Security     JWT
func (ce *EchoOrderController) GetOrderController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	order, err := ce.Service.GetOrderByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"order":    order,
	})
}

// UpdateOrderController godoc
// @Summary      Update Order Information
// @Description  User can update order information
// @Tags         Order
// @accept       json
// @Produce      json
// @Router       /order/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.Order  true  "required"
// @Success      200  {object}  model.Order
// @Failure      500   {object}  model.Order
// @Security     JWT
func (ce *EchoOrderController) UpdateOrderController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	order := model.Order{}
	c.Bind(&order)

	err := ce.Service.UpdateOrderByIDService(intID, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"order":    order,
	})
}

// DeleteOrderController godoc
// @Summary      Delete Order Information
// @Description  User can delete order information if they want it
// @Tags         Order
// @accept       json
// @Produce      json
// @Router       /order/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Order
// @Failure      500  {object}  model.Order
// @Security     JWT
func (ce *EchoOrderController) DeleteOrderController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteOrderByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
		"id":       intID,
	})
}
