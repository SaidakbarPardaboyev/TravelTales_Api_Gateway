package handler

import (
	"fmt"
	"net/http"
	pb "travel/genproto/users"

	"github.com/gin-gonic/gin"
)

// @Summary Get User information
// @Description this is for getting a user information
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id is required"
// @Success 200 {object} users.ResponseGetProfile "returns users info"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/GetProfile [get]
func (h *Handler) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("id not found from url params")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "id not found from url params",
		})
		return
	}

	req := pb.RequestGetProfile{Id: id}

	resp, err := h.UserClient.GetProfile(ctx, &req)
	if err != nil {
		h.Logger.Error("Error with requesting GetProfile method: %s", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetProfile method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
