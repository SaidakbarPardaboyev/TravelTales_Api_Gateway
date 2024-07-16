package handler

import (
	"log/slog"
	pbInter "travel/genproto/interactions"
	pbItiner "travel/genproto/itineraries"
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
	ItinerariesClient pbItiner.ItinerariesClient
}

func NewHandler() *Handler {
	userClient := connections.NewUserClient()
	storiesClient := connections.NewStoriesClient()
	interationsClient := connections.NewInteractionsClient()
	itinerariesClient := connections.NewItinerariesClient()
	logger := logger.NewLogger()
	return &Handler{
		UserClient:        userClient,
		Logger:            logger,
		StoriesClient:     storiesClient,
		InterationsClient: interationsClient,
		ItinerariesClient: itinerariesClient,
	}
}
