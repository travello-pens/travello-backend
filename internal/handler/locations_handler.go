package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoLocationController struct {
	Service adapter.AdapterLocationService
}

// CreateLocationController godoc
// @Summary      Create Location
// @Description  User can create location
// @Tags         Location
// @accept       json
// @Produce      json
// @Router       /location [post]
// @param        data  body      model.Location  true  "required"
// @Success      201   {object}  model.Location
// @Failure      500   {object}  model.Location
// @Security     JWT
func (ce *EchoLocationController) CreateLocationController(c echo.Context) error {
	location := model.Location{}
	c.Bind(&location)

	err := ce.Service.CreateLocationService(location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"location": location,
	})
}

// GetLocationsController godoc
// @Summary      Get All Locations Information
// @Description  User can get all locations information
// @Tags         Location
// @accept       json
// @Produce      json
// @Router       /location [get]
// @Success      200   {object}  model.Location
// @Security     JWT
func (ce *EchoLocationController) GetLocationsController(c echo.Context) error {
	locations := ce.Service.GetAllLocationsService()

	return c.JSONPretty(http.StatusOK, locations, " ")
}

// GetLocationController godoc
// @Summary      Get Location Information by Id
// @Description  User can get location information by id
// @Tags         Location
// @accept       json
// @Produce      json
// @Router       /location/{id} [get]
// @param        id    path      int             true  "id"
// @Success      200  {object}  model.Location
// @Failure      404  {object}  model.Location
// @Security     JWT
func (ce *EchoLocationController) GetLocationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	location, err := ce.Service.GetLocationByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, location)
}

// UpdateLocationController godoc
// @Summary      Update Location Information
// @Description  User can update location information
// @Tags         Location
// @accept       json
// @Produce      json
// @Router       /location/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.Location  true  "required"
// @Success      200  {object}  model.Location
// @Failure      500   {object}  model.Location
// @Security     JWT
func (ce *EchoLocationController) UpdateLocationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	location := model.Location{}
	c.Bind(&location)

	err := ce.Service.UpdateLocationByIDService(intID, location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"location": location,
	})
}

// DeleteLocationController godoc
// @Summary      Delete Location Information
// @Description  User can delete location information if they want it
// @Tags         Location
// @accept       json
// @Produce      json
// @Router       /location/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Location
// @Failure      500  {object}  model.Location
// @Security     JWT
func (ce *EchoLocationController) DeleteLocationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteLocationByIDService(intID)
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
