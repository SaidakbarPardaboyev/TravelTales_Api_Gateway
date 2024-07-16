package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	pb "travel/genproto/interactions"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Write Comment To Story
// @Description this is for writing comment to story
// @Tags Content
// @Accept json
// @Produce json
// @Param request body interactions.RequestCreateComment true "all params are required"
// @Success 200 {object} interactions.ResponseCreateComment "returns comment information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /interaction/comment/create [post]
func (h *Handler) CreateComment(ctx *gin.Context) {
	req := pb.RequestCreateComment{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}

	resp, err := h.InterationsClient.CreateComment(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting CreateComment method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting CreateComment method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Comments info
// @Description this is for getting comments information
// @Tags Content
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param limit query int true "limit is required"
// @Param page query int true "page is required"
// @Success 200 {object} interactions.ResponseGetComments "returns comments information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /stories/{id}/comments [get]
func (h *Handler) GetComments(ctx *gin.Context) {
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

	req := pb.RequestGetComments{
		StoryId: id,
		Limit:   int32(limit),
		Page:    int32(page),
	}

	resp, err := h.InterationsClient.GetComments(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetComments method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetComments method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// func (h *Handler) CreateComment(ctx *gin.Context) {}
