package connections

import (
	"log"
	"travel/config"
	pbInter "travel/genproto/interactions"
	pbStories "travel/genproto/stories"
	pb "travel/genproto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient() pb.UsersClient {
	conn, err := grpc.NewClient(config.Load().USER_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUsersClient(conn)
}

func NewStoriesClient() pbStories.StoriesClient {
	conn, err := grpc.NewClient(config.Load().CONTENT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return pbStories.NewStoriesClient(conn)
}

func NewInteractionsClient() pbInter.InteractionsClient {
	conn, err := grpc.NewClient(config.Load().CONTENT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return pbInter.NewInteractionsClient(conn)
}
