package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoTravelAgentController struct {
	Service adapter.AdapterTravelAgentService
}

// CreateTravelAgentController godoc
// @Summary      Create Travel Agent
// @Description  User can create agent
// @Tags         TravelAgent
// @accept       json
// @Produce      json
// @Router       /agent [post]
// @param        data  body      model.TravelAgent  true  "required"
// @Success      201   {object}  model.TravelAgent
// @Failure      500   {object}  model.TravelAgent
// @Security     JWT
func (ce *EchoTravelAgentController) CreateTravelAgentController(c echo.Context) error {
	agent := model.TravelAgent{}
	c.Bind(&agent)

	err := ce.Service.CreateTravelAgentService(agent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"agent":    agent,
	})
}

// GetTravelAgentsController godoc
// @Summary      Get All Travel Agents Information
// @Description  User can get all agents information
// @Tags         TravelAgent
// @accept       json
// @Produce      json
// @Router       /agent [get]
// @Success      200   {object}  model.TravelAgent
// @Security     JWT
func (ce *EchoTravelAgentController) GetTravelAgentsController(c echo.Context) error {
	agents := ce.Service.GetAllTravelAgentsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"agents":   agents,
	}, " ")
}

// GetTravelAgentController godoc
// @Summary      Get Travel Agent Information by Id
// @Description  User can get agent information by id
// @Tags         TravelAgent
// @accept       json
// @Produce      json
// @Router       /agent/{id} [get]
// @param        id    path      int                true  "id"
// @Success      200  {object}  model.TravelAgent
// @Failure      404  {object}  model.TravelAgent
// @Security     JWT
func (ce *EchoTravelAgentController) GetTravelAgentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	agent, err := ce.Service.GetTravelAgentByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"agent":    agent,
	})
}

// UpdateTravelAgentController godoc
// @Summary      Update Travel Agent Information
// @Description  User can update agent information
// @Tags         TravelAgent
// @accept       json
// @Produce      json
// @Router       /agent/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.TravelAgent  true  "required"
// @Success      200  {object}  model.TravelAgent
// @Failure      500   {object}  model.TravelAgent
// @Security     JWT
func (ce *EchoTravelAgentController) UpdateTravelAgentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	agent := model.TravelAgent{}
	c.Bind(&agent)

	err := ce.Service.UpdateTravelAgentByIDService(intID, agent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"agent":    agent,
	})
}

// DeleteTravelAgentController godoc
// @Summary      Delete Travel Agent Information
// @Description  User can delete agent information if they want it
// @Tags         TravelAgent
// @accept       json
// @Produce      json
// @Router       /agent/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.TravelAgent
// @Failure      500  {object}  model.TravelAgent
// @Security     JWT
func (ce *EchoTravelAgentController) DeleteTravelAgentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteTravelAgentByIDService(intID)
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
