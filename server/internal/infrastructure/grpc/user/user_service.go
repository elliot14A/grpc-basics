package user

import (
	pb "github.com/elliot14A/tcorp/assessment/proto"
	"github.com/elliot14A/tcorp/assessment/server/internal/domain/repositories"
)

type userServiceServer struct {
	userRepository repositories.UserRepository
	pb.UserServiceServer
}

func NewUserServiceServer(userRepository repositories.UserRepository) *userServiceServer {
	return &userServiceServer{userRepository: userRepository}
}
