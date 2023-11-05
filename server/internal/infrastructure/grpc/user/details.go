package user

import (
	"context"

	pb "github.com/elliot14A/tcorp/assessment/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *userServiceServer) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user, err := u.userRepository.Details(uint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	returnUser := &pb.User{
		Id:      int32(user.Id),
		Fname:   user.FirstName,
		Married: user.Married,
		City:    user.City,
		Height:  user.Height,
		Phone:   user.Phone,
	}

	return returnUser, nil
}
