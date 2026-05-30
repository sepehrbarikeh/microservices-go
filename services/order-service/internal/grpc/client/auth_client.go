package client


import (
	"context"
	"time"

	authpb "order-service/proto/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client authpb.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthClient(addr string) (*AuthClient, error) {

	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &AuthClient{
		client: authpb.NewAuthServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AuthClient) GetUserByID(userID int32) (*authpb.GetUserByIDResponse, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*3,
	)
	defer cancel()

	return c.client.GetUserByID(
		ctx,
		&authpb.GetUserByIDRequest{
			UserId: userID,
		},
	)
}