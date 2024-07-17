package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

// @Summary Get All Itineraries
// @Description this is for getting all itineraries information
// @Tags Itineraries
// @Accept json
// @Produce json
// @Param limit query int true "limit is required"
// @Param page query int true "page is required"
// @Success 200 {object} itineraries.ResponseGetAllItineraries "returns itineraries information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /itineraries [get]
func (h *Handler) GetAllItineraries(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		h.Logger.Error("limit param is invalid")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "limit param is invalid",
		})
		return
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		h.Logger.Error("page param is invalid")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "page param is invalid",
		})
		return
	}

	req := pb.RequestGetAllItineraries{
		Limit: int32(limit),
		Page:  int32(page),
	}
	resp, err := h.ItinerariesClient.GetAllItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetAllItineraries method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetAllItineraries method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Itinerary Full Info
// @Description this is for getting itinerary full information
// @Tags Itineraries
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Success 200 {object} itineraries.ResponseGetItineraryFullInfo "returns itinerary full information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /itineraries/{id}/ [get]
func (h *Handler) GetItineraryFullInfo(ctx *gin.Context) {
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

	req := pb.RequestGetItineraryFullInfo{
		Id: id,
	}

	resp, err := h.ItinerariesClient.GetItineraryFullInfo(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetItineraryFullInfo method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetItineraryFullInfo method: %s", err),
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
