package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoProductPhotoController struct {
	Service adapter.AdapterProductPhotoService
}

// CreateProductPhotoController godoc
// @Summary      Create Product Photo
// @Description  User can create product photo
// @Tags         ProductPhoto
// @accept       json
// @Produce      json
// @Router       /product-photo [post]
// @param        data  body      model.ProductPhoto  true  "required"
// @Success      201   {object}  model.ProductPhoto
// @Failure      500   {object}  model.ProductPhoto
// @Security     JWT
func (ce *EchoProductPhotoController) CreateProductPhotoController(c echo.Context) error {
	photo := model.ProductPhoto{}
	c.Bind(&photo)

	err := ce.Service.CreateProductPhotoService(photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"photo":    photo,
	})
}

// GetProductPhotosController godoc
// @Summary      Get All ProductPhotos Information
// @Description  User can get all photos information
// @Tags         ProductPhoto
// @accept       json
// @Produce      json
// @Router       /photo [get]
// @Success      200   {object}  model.ProductPhoto
// @Security     JWT
func (ce *EchoProductPhotoController) GetProductPhotosController(c echo.Context) error {
	photos := ce.Service.GetAllProductPhotosService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"photos":   photos,
	}, " ")
}

// GetProductPhotoController godoc
// @Summary      Get ProductPhoto Information by Id
// @Description  User can get photo information by id
// @Tags         ProductPhoto
// @accept       json
// @Produce      json
// @Router       /photo/{id} [get]
// @param        id    path      int                 true  "id"
// @Success      200  {object}  model.ProductPhoto
// @Failure      404  {object}  model.ProductPhoto
// @Security     JWT
func (ce *EchoProductPhotoController) GetProductPhotoController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	photo, err := ce.Service.GetProductPhotoByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"photo":    photo,
	})
}

// UpdateProductPhotoController godoc
// @Summary      Update Product Photo Information
// @Description  User can update product photo information
// @Tags         ProductPhoto
// @accept       json
// @Produce      json
// @Router       /product-photo/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.ProductPhoto  true  "required"
// @Success      200  {object}  model.ProductPhoto
// @Failure      500   {object}  model.ProductPhoto
// @Security     JWT
func (ce *EchoProductPhotoController) UpdateProductPhotoController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	photo := model.ProductPhoto{}
	c.Bind(&photo)

	err := ce.Service.UpdateProductPhotoByIDService(intID, photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"photo":    photo,
	})
}

// DeleteProductPhotoController godoc
// @Summary      Delete Product Photo Information
// @Description  User can delete product photo information if they want it
// @Tags         ProductPhoto
// @accept       json
// @Produce      json
// @Router       /product-photo/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.ProductPhoto
// @Failure      500  {object}  model.ProductPhoto
// @Security     JWT
func (ce *EchoProductPhotoController) DeleteProductPhotoController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteProductPhotoByIDService(intID)
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
