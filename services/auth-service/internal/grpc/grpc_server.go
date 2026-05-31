package server

import (
	"context"
	"log"
	"net"

	"auth-service/internal/service"
	authpb "auth-service/proto/authpb"

	"google.golang.org/grpc"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	userService *service.UserService
}

func Serve(userSvc *service.UserService, port string) {

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	authGRPCServer := NewAuthServer(userSvc)

	authpb.RegisterAuthServiceServer(grpcServer, authGRPCServer)

	log.Println("gRPC running on :", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func NewAuthServer(userService *service.UserService) *AuthServer {
	return &AuthServer{userService: userService}
}

func (s *AuthServer) GetUserByID(
	ctx context.Context,
	req *authpb.GetUserByIDRequest,
) (*authpb.GetUserByIDResponse, error) {

	user, err := s.userService.GetUserByID(int(req.UserId))
	if err != nil {
		return nil, err
	}

	return &authpb.GetUserByIDResponse{
		Id:    int32(user.ID),
		Email: user.Email,
	}, nil
}

