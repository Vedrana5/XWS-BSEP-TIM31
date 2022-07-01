package mapper

import (
	pb "common/module/proto/user_service"
	"user/module/dto"
)

func MapUser(user dto.RegisteredUserDTO) *pb.RegisterUser {
	usersPb := &pb.RegisterUser{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return usersPb
}
