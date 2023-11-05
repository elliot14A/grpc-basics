package grpc

import (
	"fmt"
	"net"

	pb "github.com/elliot14A/tcorp/assessment/proto"
	"github.com/elliot14A/tcorp/assessment/server/app"
	"github.com/elliot14A/tcorp/assessment/server/internal/domain/repositories"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/grpc/user"
	"google.golang.org/grpc"
)

func RunGrpcServer(a *app.App) {
	logger := a.GetLogger()
	database := a.GetDatabase()
	config := a.GetConfig()

	lis, err := net.Listen("tcp", ":"+config.GetServerConfig().ServerPort)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error starting tcp server: %v", err.Error()))
	}

	grpcServer := grpc.NewServer()
	userRepository, err := repositories.NewUserRepository(database)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error creating repository: %v", err.Error()))
	}
	userServiceServer := user.NewUserServiceServer(userRepository)
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)
	logger.Info(fmt.Sprintf("server started at port: %v", config.GetServerConfig().ServerPort))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve: %s", err)
	}
}
