package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	pb "travel/genproto/users"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// @Summary Get User information
// @Description this is for getting a user information
// @Tags users
// @Accept json
// @Produce json
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
		h.Logger.Error(fmt.Sprintf("error with requesting GetProfile method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetProfile method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Validate User
// @Description this is for chacking user is exists
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Success 200 {object} users.Status "returns boolean value about existing user"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/validate [get]
func (h *Handler) ValidateUser(ctx *gin.Context) {
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
	req := pb.RequestGetProfile{Id: id}

	resp, err := h.UserClient.ValidateUser(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting ValidateUser method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting ValidateUser method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Edit Profile
// @Description this is for editing user information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body users.RequestEditProfile true "all params are required"
// @Success 200 {object} users.ResponseEditProfile "returns user information about existing user"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/editprofile [post]
func (h *Handler) EditProfile(ctx *gin.Context) {
	req := pb.RequestEditProfile{}

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

	resp, err := h.UserClient.EditProfile(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting EditProfile method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting EditProfile method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Users info
// @Description this is for getting user information
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int true "limit is required"
// @Param page query int true "page is required"
// @Success 200 {object} users.ResponseGetUsers "returns users information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users [get]
func (h *Handler) GetUsers(ctx *gin.Context) {
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

	req := pb.RequestGetUsers{
		Limit: int32(limit),
		Page:  int32(page),
	}
	resp, err := h.UserClient.GetUsers(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetUsers method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetUsers method: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete User
// @Description This endpoint is for deleting a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} users.ResponseDeleteUser "Message about deleting user"
// @Failure 400 {object} models.Error "Occurs when user enters invalid params"
// @Failure 500 {object} models.Error "Occurs when an internal service error happens"
// @Router /users/{id}/delete [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("id not found from url params")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "id not found from url params",
		})
		return
	}

	req := pb.RequestDeleteUser{Id: id}

	resp, err := h.UserClient.DeleteUser(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting DeleteUser method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting DeleteUser method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Password
// @Description This endpoint is for updating password
// @Tags users
// @Accept json
// @Produce json
// @Param request body users.RequestUpdatePassword true "User ID"
// @Success 200 {object} users.ResponseUpdatePassword "Message about updating password"
// @Failure 400 {object} models.Error "Occurs when user enters invalid params"
// @Failure 500 {object} models.Error "Occurs when an internal service error happens"
// @Router /users/update [put]
func (h *Handler) UpdatePassword(ctx *gin.Context) {
	req := pb.RequestUpdatePassword{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with decoding url body: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": fmt.Sprintf("Error with decoding url body: %s", err.Error()),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error with hashing password: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with hashing password: %s", err.Error()),
		})
		return
	}
	req.NewPassword = string(hashedPassword)

	resp, err := h.UserClient.UpdatePassword(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting UpdatePassword method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting UpdatePassword method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Follow
// @Description this is for following user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param request body users.RequestFollow true "all params are required and followoreId is taking from params"
// @Success 200 {object} users.ResponseFollow "returns all ditails about following"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/follow [post]
func (h *Handler) Follow(ctx *gin.Context) {
	req := pb.RequestFollow{}

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
	req.FollowerId = id

	resp, err := h.UserClient.Follow(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting Follow method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting Follow method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Followers
// @Description this is for getting followers information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id is required"
// @Param limit query int true "limit is required"
// @Param page query int true "page is required"
// @Success 200 {object} users.ResponseGetFollowers "returns followers information"
// @Failure 400 {object} models.Error "It occurs when user enter invalid params"
// @Failure 500 {object} models.Error "It occurs when error happenes internal service"
// @Router /users/{id}/followers [get]
func (h *Handler) GetFollowers(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("id not found from url params")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "http.StatusBadRequest",
			"massege": "id not found from url params",
		})
		return
	}

	req := pb.RequestGetFollowers{UserId: id}

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

	req.Limit = int32(limit)
	req.Page = int32(page)

	resp, err := h.UserClient.GetFollowers(ctx, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error with requesting GetFollowers method: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "http.StatusInternalServerError",
			"massege": fmt.Sprintf("Error with requesting GetFollowers method: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
// func (h *Handler) ValidateUser(ctx *gin.Context) {}
