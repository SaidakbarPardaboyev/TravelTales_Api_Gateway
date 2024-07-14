package handler

import (
	"log/slog"
	pb "travel/genproto/users"
	"travel/pkg/connections"
	"travel/pkg/logger"
)

type Handler struct {
	Logger     *slog.Logger
	UserClient pb.UsersClient
}

func NewHandler() *Handler {
	userClient := connections.NewUserClient()
	logger := logger.NewLogger()
	return &Handler{
		UserClient: userClient,
		Logger:     logger,
	}
}
