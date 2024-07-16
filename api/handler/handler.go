package handler

import (
	"log/slog"
	pb "travel/genproto/users"
	pbStories "travel/genproto/stories"
	
	"travel/pkg/connections"
	"travel/pkg/logger"
)

type Handler struct {
	Logger     *slog.Logger
	UserClient pb.UsersClient
	StoriesClient pbStories.StoriesClient
}

func NewHandler() *Handler {
	userClient := connections.NewUserClient()
	storiesClient := connections.NewStoriesClient()
	logger := logger.NewLogger()
	return &Handler{
		UserClient: userClient,
		Logger:     logger,
		StoriesClient: storiesClient,
	}
}
