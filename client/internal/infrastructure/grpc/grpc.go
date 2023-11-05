package grpc

import (
	"context"
	"fmt"
	"io"

	"github.com/elliot14A/tcorp/assessment/client/internal/domain/models"
	"github.com/elliot14A/tcorp/assessment/pkg/config"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
	pb "github.com/elliot14A/tcorp/assessment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcUserRepository struct {
	client pb.UserServiceClient
}

func NewGrpcUserRepository(config config.ServerConfig, logger logger.LoggerPort) *GrpcUserRepository {
	conn, err := grpc.Dial("localhost"+":"+config.ServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot connect to server: %v", err.Error()))
	}

	client := pb.NewUserServiceClient(conn)

	return &GrpcUserRepository{client: client}
}

func (g *GrpcUserRepository) Details(userId int) (*models.User, error) {
	user, err := g.client.GetUserById(context.Background(), &pb.UserId{Id: int32(userId)})
	if err != nil {
		return nil, err
	}
	return &models.User{
		Id:        uint(user.Id),
		FirstName: user.Fname,
		Married:   user.Married,
		City:      user.City,
		Phone:     user.Phone,
		Height:    user.Height,
	}, nil
}

func (g *GrpcUserRepository) List() ([]models.User, error) {
	stream, err := g.client.GetAllUsers(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	var users []models.User

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		users = append(users, models.User{
			Id:        uint(user.Id),
			FirstName: user.Fname,
			Married:   user.Married,
			City:      user.City,
			Phone:     user.Phone,
			Height:    user.Height,
		})
	}
	return users, nil
}
