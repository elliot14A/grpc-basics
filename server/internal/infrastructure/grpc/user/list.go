package user

import pb "github.com/elliot14A/tcorp/assessment/proto"

func (u *userServiceServer) GetAllUsers(req *pb.Empty, stream pb.UserService_GetAllUsersServer) error {
	users := u.userRepository.List()
	for _, user := range users {
		stream.Send(&pb.User{
			Id:      int32(user.Id),
			Fname:   user.FirstName,
			Married: user.Married,
			City:    user.City,
			Height:  user.Height,
			Phone:   user.Phone,
		})
	}
	return nil
}
