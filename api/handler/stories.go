package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	pb "travel/genproto/stories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create Story
// @Description this is for creating story
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body stories.RequestCreateStory true "all params are required"
// @Success 200 {object} stories.ResponseCreateStory "returns story information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/stories/create [post]
func (h *Handler) CreateStory(ctx *gin.Context) {
	req := pb.RequestCreateStory{}

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
	req.AuthorId = id

	resp, err := h.StoriesClient.CreateStory(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting CreateStory method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting CreateStory method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Edit Story
// @Description this is for editing story
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body stories.RequestEditStory true "all params are required"
// @Success 200 {object} stories.ResponseEditStory "returns story information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /stories/{id}/edit [put]
func (h *Handler) EditStory(ctx *gin.Context) {
	req := pb.RequestEditStory{}

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

	resp, err := h.StoriesClient.EditStory(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting EditStory method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting EditStory method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get stories info
// @Description this is for getting stories information
// @Tags Stories
// @Accept json
// @Produce json
// @Param limit query int true "limit is required"
// @Param page query int true "page is required"
// @Success 200 {object} stories.ResponseGetStories "returns stories information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /stories [get]
func (h *Handler) GetStories(ctx *gin.Context) {
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

	req := pb.RequestGetStories{
		Limit: int32(limit),
		Page:  int32(page),
	}
	resp, err := h.StoriesClient.GetStories(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetStories method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetStories method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Full Story Info
// @Description this is for getting full information about a story
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Success 200 {object} stories.ResponseGetStoryFullInfo "returns full story information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /stories/{id}/fullinfo [get]
func (h *Handler) GetStoryFullInfo(ctx *gin.Context) {
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
	req := pb.RequestGetStoryFullInfo{Id: id}

	resp, err := h.StoriesClient.GetStoryFullInfo(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetStoryFullInfo method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetStoryFullInfo method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Story
// @Description this is for deleting story
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Success 200 {object} stories.ResponseDeleteStory "returns message about deleting story"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /stories/{id}/delete [delete]
func (h *Handler) DeleteStory(ctx *gin.Context) {
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
	req := pb.RequestDeleteStory{StoryId: id}

	resp, err := h.StoriesClient.DeleteStory(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting DeleteStory method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting DeleteStory method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
