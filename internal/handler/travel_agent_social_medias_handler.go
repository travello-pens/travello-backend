package handler

import (
	"net/http"
	"strconv"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoTravelAgentSocialMediaController struct {
	Service adapter.AdapterTravelAgentSocialMediaService
}

// CreateTravelAgentSocialMediaController godoc
// @Summary      Create Travel Agent Social Media
// @Description  User can create travel agent social media information
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media [post]
// @param        data  body      model.TravelAgentSocialMedia  true  "required"
// @Success      201   {object}  model.TravelAgentSocialMedia
// @Failure      500   {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) CreateTravelAgentSocialMediaController(c echo.Context) error {
	media := model.TravelAgentSocialMedia{}
	c.Bind(&media)

	err := ce.Service.CreateTravelAgentSocialMediaService(media)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"media":    media,
	})
}

// GetTravelAgentSocialMediasController godoc
// @Summary      Get All Travel Agents Social Media Information
// @Description  User can get all travel agents social media information
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media [get]
// @Success      200   {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) GetTravelAgentSocialMediasController(c echo.Context) error {
	medias := ce.Service.GetAllTravelAgentSocialMediasService()

	return c.JSONPretty(http.StatusOK, medias, " ")
}

// GetTravelAgentSocialMediaController godoc
// @Summary      Get Travel Agent Social Media Information by Id
// @Description  User can get travel agent social media information by id
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media/{id} [get]
// @param        id    path      int                           true  "id"
// @Success      200  {object}  model.TravelAgentSocialMedia
// @Failure      404  {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) GetTravelAgentSocialMediaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	media, err := ce.Service.GetTravelAgentSocialMediaByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, media)
}

// GetTravelAgentSocialMediaByTravelAgentController godoc
// @Summary      Get Travel Agent Social Media Information by Travel Agent
// @Description  User can get travel agent social media information by travel agent
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media/agent/{agent} [get]
// @param        agent    path      string            true  "agent"
// @Success      200  {object}  model.TravelAgentSocialMedia
// @Failure      404  {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) GetTravelAgentSocialMediaByTravelAgentController(c echo.Context) error {
	agent := c.Param("agent")

	media, err := ce.Service.GetTravelAgentSocialMediaByTravelAgentService(agent)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no agent travel",
		})
	}

	return c.JSON(http.StatusOK, media)
}

// UpdateTravelAgentSocialMediaController godoc
// @Summary      Update Travel Agent Social Media Information
// @Description  User can update travel agent social media information
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.TravelAgentSocialMedia  true  "required"
// @Success      200  {object}  model.TravelAgentSocialMedia
// @Failure      500   {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) UpdateTravelAgentSocialMediaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	media := model.TravelAgentSocialMedia{}
	c.Bind(&media)

	err := ce.Service.UpdateTravelAgentSocialMediaByIDService(intID, media)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
		"media":    media,
	})
}

// DeleteTravelAgentSocialMediaController godoc
// @Summary      Delete Travel Agent Social Media Information
// @Description  User can delete travel agent social media information if they want it
// @Tags         TravelAgentSocialMedia
// @accept       json
// @Produce      json
// @Router       /travel-agent-social-media/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.TravelAgentSocialMedia
// @Failure      500  {object}  model.TravelAgentSocialMedia
// @Security     JWT
func (ce *EchoTravelAgentSocialMediaController) DeleteTravelAgentSocialMediaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteTravelAgentSocialMediaByIDService(intID)
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
