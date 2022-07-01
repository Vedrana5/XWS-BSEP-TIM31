package mapper

import (
	pb "common/module/proto/user_service"
	"user/module/dto"
	"user/module/model"
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

func MapFindPublicUser(user model.User) *pb.User {
	usersPb := &pb.User{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return usersPb
}

func MapString(editUser dto.EditProfileDTO) *pb.EditUser1 {
	message := &pb.EditUser1{
		ID:             editUser.ID.String(),
		OldUsername:    editUser.OldUsername,
		Username:       editUser.Username,
		Password:       editUser.Password,
		Email:          editUser.Email,
		PhoneNumber:    editUser.PhoneNumber,
		DateOfBirth:    editUser.DateOfBirth,
		FirstName:      editUser.FirstName,
		LastName:       editUser.LastName,
		Gender:         editUser.Gender,
		TypeOfUser:     editUser.TypeOfUser,
		TypeOfProfile:  editUser.TypeOfProfile,
		Biography:      editUser.Biography,
		WorkExperience: editUser.WorkExperience,
		Education:      editUser.Education,
		Skills:         editUser.Skills,
		Interest:       editUser.Interest,
	}
	return message
}

func MapLoginUser(user dto.LogInResponseDTO) *pb.LoginUserResponse {
	var typee = "ADMIN"
	if user.TypeOfUser == 0 {
		typee = "ADMIN"
	}
	if user.TypeOfUser == 1 {
		typee = "REGISTERED_USER"
	}
	if user.TypeOfUser == 2 {
		typee = "UNREGISTERED_USER"
	}
	usersPb := &pb.LoginUserResponse{
		ID:         user.ID.String(),
		Token:      user.Token,
		TypeOfUser: typee,
	}
	return usersPb
}

func MapFindUser(user *model.User) *pb.User {
	var typeOfUser = "ADMIN"
	if user.TypeOfUser == 0 {
		typeOfUser = "ADMIN"
	}
	if user.TypeOfUser == 1 {
		typeOfUser = "REGISTERED_USER"
	}
	if user.TypeOfUser == 2 {
		typeOfUser = "UNREGISTERED_USER"
	}
	var typeOfProfile = "PUBLIC"
	if user.TypeOfProfile == 1 {
		typeOfProfile = "PRIVATE"
	}
	var gender = "MALE"
	if user.Gender == 1 {
		gender = "FEMALE"
	}
	if user.Gender == 2 {
		gender = "OTHER"
	}
	usersPb := &pb.User{
		Username:       user.Username,
		Password:       user.Password,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		DateOfBirth:    user.DateOfBirth.String(),
		TypeOfUser:     typeOfUser,
		TypeOfProfile:  typeOfProfile,
		Gender:         gender,
		Biography:      user.Biography,
		WorkExperience: user.WorkExperience,
		Education:      user.Education,
		Skills:         user.Skills,
		Interest:       user.Interest,
		Question:       user.Question,
		Answer:         user.Answer,
	}
	return usersPb
}

/*func MapFindPublicUser(user []*model.User) []pb.User {
	var usersPb []pb.User
	for i := 0; i < len(user); i++ {
		var typeOfUser = "ADMIN"
		if user[i].TypeOfUser == 0 {
			typeOfUser = "ADMIN"
		}
		if user[i].TypeOfUser == 1 {
			typeOfUser = "REGISTERED_USER"
		}
		if user[i].TypeOfUser == 2 {
			typeOfUser = "UNREGISTERED_USER"
		}
		var typeOfProfile = "PUBLIC"
		if user[i].TypeOfProfile == 1 {
			typeOfProfile = "PRIVATE"
		}
		var gender = "MALE"
		if user[i].Gender == 1 {
			gender = "FEMALE"
		}
		if user[i].Gender == 2 {
			gender = "OTHER"
		}
		var userPb pb.User
		userPb = pb.User{
			Username:       user[i].Username,
			Password:       user[i].Password,
			Email:          user[i].Email,
			PhoneNumber:    user[i].PhoneNumber,
			FirstName:      user[i].FirstName,
			LastName:       user[i].LastName,
			DateOfBirth:    user[i].DateOfBirth.String(),
			TypeOfUser:     typeOfUser,
			TypeOfProfile:  typeOfProfile,
			Gender:         gender,
			Biography:      user[i].Biography,
			WorkExperience: user[i].WorkExperience,
			Education:      user[i].Education,
			Skills:         user[i].Skills,
			Interest:       user[i].Interest,
			Question:       user[i].Question,
			Answer:         user[i].Answer,
		}

		usersPb = append(usersPb, userPb)

	}

	return usersPb
}*/
