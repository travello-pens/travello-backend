package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoProductController struct {
	Service adapter.AdapterProductService
}

// CreateProductController godoc
// @Summary      Create Product
// @Description  User can create product
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product [post]
// @param        data  body      model.Product  true  "required"
// @Success      201   {object}  model.Product
// @Failure      500   {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) CreateProductController(c echo.Context) error {
	product := model.Product{}
	c.Bind(&product)

	err := ce.Service.CreateProductService(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"product":  product,
	})
}

// GetProductsController godoc
// @Summary      Get All Products Information
// @Description  User can get all products information
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product [get]
// @Success      200   {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) GetProductsController(c echo.Context) error {
	products := ce.Service.GetAllProductsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"products": products,
	}, " ")
}

// GetProductController godoc
// @Summary      Get Product Information by Id
// @Description  User can get product information by id
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product/{id} [get]
// @param        id    path      int            true  "id"
// @Success      200  {object}  model.Product
// @Failure      404  {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) GetProductController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	product, err := ce.Service.GetProductByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"product":  product,
	})
}

// GetProductsByLocationController godoc
// @Summary      Get Product Information by Location
// @Description  User can get product information by location
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product/location/{location} [get]
// @param        location    path      string            true  "location"
// @Success      200  {object}  model.Product
// @Failure      404  {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) GetProductsByLocationController(c echo.Context) error {
	location := c.Param("location")

	products, err := ce.Service.GetProductByLocationService(location)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no location",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"products": products,
	})
}

// GetProductsByAgentController godoc
// @Summary      Get Product Information by Travel Agent
// @Description  User can get product information by travel agent
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product/agent/{agent} [get]
// @param        agent    path      string            true  "agent"
// @Success      200  {object}  model.Product
// @Failure      404  {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) GetProductsByAgentController(c echo.Context) error {
	agent := c.Param("agent")

	products, err := ce.Service.GetProductByAgentService(agent)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no agent travel",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"products": products,
	})
}

// UpdateProductController godoc
// @Summary      Update Product Information
// @Description  User can update product information
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.Product  true  "required"
// @Success      200  {object}  model.Product
// @Failure      500   {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) UpdateProductController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	product := model.Product{}
	c.Bind(&product)

	err := ce.Service.UpdateProductByIDService(intID, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"product":  product,
	})
}

// DeleteProductController godoc
// @Summary      Delete Product Information
// @Description  User can delete product information if they want it
// @Tags         Product
// @accept       json
// @Produce      json
// @Router       /product/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Product
// @Failure      500  {object}  model.Product
// @Security     JWT
func (ce *EchoProductController) DeleteProductController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteProductByIDService(intID)
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
