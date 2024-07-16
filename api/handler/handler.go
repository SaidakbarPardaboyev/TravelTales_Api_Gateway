package handler

import (
	"log/slog"
	pbInter "travel/genproto/interactions"
	pbStories "travel/genproto/stories"
	pb "travel/genproto/users"

	"travel/pkg/connections"
	"travel/pkg/logger"
)

type Handler struct {
	Logger            *slog.Logger
	UserClient        pb.UsersClient
	StoriesClient     pbStories.StoriesClient
	InterationsClient pbInter.InteractionsClient
}

func NewHandler() *Handler {
	userClient := connections.NewUserClient()
	storiesClient := connections.NewStoriesClient()
	interationsClient := connections.NewInteractionsClient()
	logger := logger.NewLogger()
	return &Handler{
		UserClient:        userClient,
		Logger:            logger,
		StoriesClient:     storiesClient,
		InterationsClient: interationsClient,
	}
}
