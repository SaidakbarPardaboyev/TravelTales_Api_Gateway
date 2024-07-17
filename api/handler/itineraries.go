package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	pb "travel/genproto/itineraries"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create Itineraries
// @Description this is for creating itineraries
// @Tags Itineraries
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body itineraries.RequestCreateItineraries true "all params are required"
// @Success 200 {object} itineraries.ResponseCreateItineraries "returns itinerary information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/itineraries/create [post]
func (h *Handler) CreateItineraries(ctx *gin.Context) {
	req := pb.RequestCreateItineraries{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("id not found from url params")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "id not found from url params",
		})
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Error("invalid uuid")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "invalid uuid",
		})
		return
	}

	req.AutherId = id

	resp, err := h.ItinerariesClient.CreateItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting CreateItineraries method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting CreateItineraries method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Edit Itinerary
// @Description this is for editing itinerary
// @Tags Itineraries
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body itineraries.RequestEditItineraries true "all params are required"
// @Success 200 {object} itineraries.ResponseEditItineraries "returns itinerary information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /itineraries/{id}/edit [put]
func (h *Handler) EditItineraries(ctx *gin.Context) {
	req := pb.RequestEditItineraries{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("id not found from url params")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "id not found from url params",
		})
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Error("invalid uuid")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "invalid uuid",
		})
		return
	}

	req.Id = id

	resp, err := h.ItinerariesClient.EditItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting EditItineraries method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting EditItineraries method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
// func (h *Handler) CreateItineraries(ctx *gin.Context) {}
