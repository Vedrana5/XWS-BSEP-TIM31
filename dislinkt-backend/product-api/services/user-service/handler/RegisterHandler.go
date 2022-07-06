package handler

import (
	"common/module/proto/user_service"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"log"
	"net/smtp"
	"strings"
	"time"
	"user/module/dto"
	"user/module/mapper"
	"user/module/model"
	"user/module/service"
	"user/module/util"
)

type RegisterHandler struct {
	Rbac                     *gorbac.RBAC
	PermissionUpdateUserInfo *gorbac.Permission
	ConfirmationTokenService *service.ConfirmationTokenService
	PasswordUtil             *util.PasswordUtil
	UserService              *service.UserService
	LogInfo                  *logrus.Logger
	LogError                 *logrus.Logger
}

func (u RegisterHandler) MustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewRegisterHandler(rbac *gorbac.RBAC, PermissionUpdateUserInfo *gorbac.Permission, confirmationTokenService *service.ConfirmationTokenService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, userService *service.UserService) *RegisterHandler {
	return &RegisterHandler{rbac, PermissionUpdateUserInfo, confirmationTokenService, passwordUtil, userService, LogInfo, LogError}
}

func (handler RegisterHandler) CreateUser(ctx context.Context, user *user_service.CreateRequest) (*user_service.CreateResponse, error) {

	var registeredUserDTO dto.RegisteredUserDTO
	registeredUserDTO = dto.RegisteredUserDTO{
		Username:       user.User.Username,
		Password:       user.User.Password,
		Email:          user.User.Email,
		PhoneNumber:    user.User.PhoneNumber,
		FirstName:      user.User.FirstName,
		LastName:       user.User.LastName,
		DateOfBirth:    user.User.DateOfBirth,
		TypeOfUser:     user.User.TypeOfUser,
		TypeOfProfile:  user.User.TypeOfProfile,
		Gender:         user.User.Gender,
		Biography:      user.User.Biography,
		WorkExperience: user.User.WorkExperience,
		Education:      user.User.Education,
		Skills:         user.User.Skills,
		Interest:       user.User.Interest,
		Question:       user.User.Question,
		Answer:         user.User.Answer,
	}

	var registeredUserResponseDTO dto.RegisteredUserDTO
	registeredUserResponseDTO = dto.RegisteredUserDTO{
		Username:  user.User.Username,
		Email:     user.User.Email,
		FirstName: user.User.FirstName,
		LastName:  user.User.LastName,
	}

	if handler.UserService.FindByUserName(registeredUserDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered username!")
		fmt.Println(time.Now().String() + " User already exist with entered username!")
		return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
	}

	if handler.UserService.FindByEmail(registeredUserDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered email!")
		fmt.Println(time.Now().String() + " User already exist with entered email!")
		return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
	}

	salt := ""
	password := ""

	answer := ""
	answerSalt := ""

	validPassword := handler.PasswordUtil.IsValidPassword(registeredUserDTO.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(registeredUserDTO.Password)

		//ANSWER SALT
		answerSalt, answer = handler.PasswordUtil.GeneratePasswordWithSalt(registeredUserDTO.Answer)

	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		fmt.Println(time.Now().String() + " Password doesn't in valid format!")

		return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
	}

	log.Printf(registeredUserDTO.Gender)
	gender := model.OTHER
	switch registeredUserDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	var typeofUser = model.REGISTERED_USER
	log.Printf(registeredUserDTO.TypeOfProfile)
	typeOfProfile := model.PRIVATE
	switch registeredUserDTO.TypeOfProfile {
	case "PUBLIC":
		typeOfProfile = model.PUBLIC
	case "PRIVATE":
		typeOfProfile = model.PRIVATE
	}

	userId := uuid.New()
	layout := "2006-01-02"
	dateOfBirth, _ := time.Parse(layout, registeredUserDTO.DateOfBirth)
	registeredUser := model.User{
		ID:             userId,
		Username:       registeredUserDTO.Username,
		Password:       password,
		Email:          registeredUserDTO.Email,
		PhoneNumber:    registeredUserDTO.PhoneNumber,
		FirstName:      registeredUserDTO.FirstName,
		LastName:       registeredUserDTO.LastName,
		DateOfBirth:    dateOfBirth,
		Salt:           salt,
		TypeOfUser:     typeofUser,
		TypeOfProfile:  typeOfProfile,
		Gender:         gender,
		Biography:      registeredUserDTO.Biography,
		WorkExperience: registeredUserDTO.WorkExperience,
		Education:      registeredUserDTO.Education,
		Skills:         registeredUserDTO.Skills,
		Interest:       registeredUserDTO.Interest,
		Question:       registeredUserDTO.Question,
		Answer:         answer,
		AnswerSalt:     answerSalt,
		IsConfirmed:    false,
	}
	confirmationToken := model.ConfirmationToken{
		ID:                uuid.New(),
		ConfirmationToken: uuid.New(),
		UserId:            userId,
		CreatedTime:       time.Now(),
		ExpiredTime:       time.Now().Add(time.Minute * 120),
		IsValid:           true,
	}

	if err := handler.ConfirmationTokenService.CreateConfirmationToken(&confirmationToken); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("Failed creating confirmation token for user!")
		fmt.Println(time.Now().String() + " Failed creating confirmation token for user!")

		return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
	}

	handler.SendConfirmationMail(registeredUser, confirmationToken.ConfirmationToken)

	if err := handler.UserService.CreateUser(&registeredUser); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("Failed creating basic user!")
		fmt.Println(time.Now().String() + " Failed creating basic user!")
		return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
	}
	return &user_service.CreateResponse{RegisterUser: mapper.MapUser(registeredUserResponseDTO)}, nil
}

func (handler RegisterHandler) LoginUser(ctx context.Context, loginUser *user_service.LoginRequest) (*user_service.LoginResponse, error) {
	var logInUserDTO dto.LogInUserDTO
	logInUserDTO = dto.LogInUserDTO{
		Username: loginUser.LoginUser.Username,
		Password: loginUser.LoginUser.Password,
	}

	var user = handler.UserService.FindByUserName(logInUserDTO.Username)

	logInResponse1 := dto.LogInResponseDTO{
		ID:         user.ID,
		Token:      "",
		TypeOfUser: user.TypeOfUser,
	}
	fmt.Print("TIP USERA JE je ", logInResponse1.TypeOfUser)
	validPassword := handler.PasswordUtil.IsValidPassword(logInUserDTO.Password)
	plainPassword := ""
	if !validPassword {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogIn",
			"timestamp": time.Now().String(),
		}).Error("Password isn't in valid format!")
		fmt.Println(time.Now().String() + " Password isn't in valid format!")
		return &user_service.LoginResponse{LoginUserResponse: mapper.MapLoginUser(logInResponse1)}, nil
	} else {
		var sb strings.Builder
		salt := user.Salt
		sb.WriteString(logInUserDTO.Password)
		sb.WriteString(salt)
		plainPassword = sb.String()
	}

	if !handler.PasswordUtil.CheckPasswordHash(plainPassword, user.Password) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogIn",
			"timestamp": time.Now().String(),
		}).Error("Failed sign up!")
		fmt.Println(time.Now().String() + " Failed sign up!")
		return &user_service.LoginResponse{LoginUserResponse: mapper.MapLoginUser(logInResponse1)}, nil
	}

	//token
	token, err := CreateToken(user.Username)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogIn",
			"timestamp": time.Now().String(),
		}).Error("Failed creating AWT token!")
		fmt.Println(time.Now().String() + " Failed creating AWT token!")

		return &user_service.LoginResponse{LoginUserResponse: mapper.MapLoginUser(logInResponse1)}, nil
	}

	logInResponse := dto.LogInResponseDTO{
		ID:         user.ID,
		Token:      token,
		TypeOfUser: user.TypeOfUser,
	}
	fmt.Print("Response je ", logInResponse)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "LogInHandler",
		"action":    "LogIn",
		"timestamp": time.Now().String(),
	}).Info("Successfully sign in user")
	fmt.Println(time.Now().String() + " Successfully sign in user!")

	return &user_service.LoginResponse{LoginUserResponse: mapper.MapLoginUser(logInResponse)}, nil
}

//SendConfirmationMail
func (handler RegisterHandler) SendConfirmationMail(user model.User, token uuid.UUID) {
	// Sender data.
	from := "sammilica99@gmail.com"
	password := "setmkiwpicaxhmti"

	// Receiver email address.
	to := []string{
		user.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Dear " + user.FirstName + ",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttps://localhost:8082/confirmRegistration/" + token.String() + "/" + user.ID.String())

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisterHandler",
		"action":    "SendConfirmationMail",
		"timestamp": time.Now().String(),
	}).Info("Successfully sended email with confirmation token!")
	fmt.Println(time.Now().String() + " Successfully sended email with confirmation token!")

}

func (handler RegisterHandler) FindByUsername(ctx context.Context, username *user_service.UserNameRequest) (*user_service.UserNameResponse, error) {

	var user = handler.UserService.FindByUserName(username.Username)
	if user == nil {
		var user1 = model.User{
			Username:       "",
			Password:       "",
			Email:          "",
			PhoneNumber:    "",
			FirstName:      "",
			LastName:       "",
			DateOfBirth:    time.Now(),
			Gender:         model.MALE,
			TypeOfUser:     model.REGISTERED_USER,
			TypeOfProfile:  model.PUBLIC,
			Biography:      "",
			WorkExperience: "",
			Education:      "",
			Skills:         "",
			Interest:       "",
			Question:       "",
			Answer:         "",
		}
		return &user_service.UserNameResponse{User: mapper.MapFindUser1(user1)}, nil
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindByUserName",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")

	return &user_service.UserNameResponse{User: mapper.MapFindUser(user)}, nil
}

func (handler RegisterHandler) EditUser(ctx context.Context, editUser *user_service.EditRequest) (*user_service.EditResponse, error) {

	id := uuid.MustParse(editUser.EditUser.ID)
	var EditProfileDTO dto.EditProfileDTO
	EditProfileDTO = dto.EditProfileDTO{
		ID:             id,
		OldUsername:    editUser.EditUser.OldUsername,
		Username:       editUser.EditUser.Username,
		Password:       editUser.EditUser.Password,
		Email:          editUser.EditUser.Email,
		PhoneNumber:    editUser.EditUser.PhoneNumber,
		DateOfBirth:    editUser.EditUser.DateOfBirth,
		FirstName:      editUser.EditUser.FirstName,
		LastName:       editUser.EditUser.LastName,
		Gender:         editUser.EditUser.Gender,
		TypeOfUser:     editUser.EditUser.TypeOfUser,
		TypeOfProfile:  editUser.EditUser.TypeOfProfile,
		Biography:      editUser.EditUser.Biography,
		WorkExperience: editUser.EditUser.WorkExperience,
		Education:      editUser.EditUser.Education,
		Skills:         editUser.EditUser.Skills,
		Interest:       editUser.EditUser.Interest,
	}

	fmt.Printf("BIOGRAFIJA JE" + EditProfileDTO.Biography)
	var loginUser = handler.UserService.FindByID(EditProfileDTO.ID)
	userRole := ""
	if loginUser.TypeOfUser == model.ADMIN {
		userRole = "role-admin"
	} else if loginUser.TypeOfUser == model.REGISTERED_USER {
		userRole = "role-registered-user"
	} else {
		userRole = "role-unregistered-user"
	}
	log.Print("ROLA JE" + userRole)
	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("User is not authorized to update user information!")
		fmt.Println(time.Now().String() + " User is not authorized to update user information!")

		return &user_service.EditResponse{EditUser1: mapper.MapString(EditProfileDTO)}, nil
	}

	if handler.UserService.FindByEmailAndID(EditProfileDTO.ID, EditProfileDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered email!")
		fmt.Println(time.Now().String() + " User already exist with entered email!")

		return &user_service.EditResponse{EditUser1: mapper.MapString(EditProfileDTO)}, nil
	}

	fmt.Println("ID JE " + EditProfileDTO.ID.String())
	fmt.Println("USERNAME JE" + EditProfileDTO.Username)

	if handler.UserService.FindByUserNameAndID(EditProfileDTO.ID, EditProfileDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered username!")
		fmt.Println(time.Now().String() + " User already exist with entered username!")

		return &user_service.EditResponse{EditUser1: mapper.MapString(EditProfileDTO)}, nil
	}

	err := handler.UserService.UpdateUserProfileInfo(&EditProfileDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("Failed updating basic user profile information!")
		fmt.Println(time.Now().String() + " Failed updating basic user profile information!")
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UpdateProfileHandler",
		"action":    "UpdateUserProfileInfo",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated user profile info!")
	fmt.Println(time.Now().String() + " Successfully updated user profile info!")

	return &user_service.EditResponse{EditUser1: mapper.MapString(EditProfileDTO)}, nil
}

func (handler RegisterHandler) FindPublicByUsername(ctx context.Context, username *user_service.PublicUserNameRequest) (*user_service.PublicUserNameResponse, error) {
	fmt.Printf("Ime koje trazi je" + username.Username)
	var users = handler.UserService.FindPublic(username.Username)
	if users == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindByUserName",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " Users not found!")
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindByUserName",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")

	response := &user_service.PublicUserNameResponse{
		Users: []*user_service.User{},
	}
	fmt.Printf("Pre kreiranja response-a")
	for _, User := range users {
		current := mapper.MapFindPublicUser(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler RegisterHandler) ChangePassword(ctx context.Context, changePassword *user_service.ChangePasswordRequest) (*user_service.ChangePasswordResponse, error) {
	fmt.Printf("USLA SAM tu i id je" + changePassword.ChangePassword.ID)
	id := uuid.MustParse(changePassword.ChangePassword.ID)
	fmt.Printf("ISPARSIRALA SAM")
	var requestDTO dto.RequestDTO
	requestDTO = dto.RequestDTO{
		ID:       id,
		Password: changePassword.ChangePassword.Password,
		Token:    changePassword.ChangePassword.Token,
	}

	var user = handler.UserService.FindByID(requestDTO.ID)

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("User is not found!")
		fmt.Println(time.Now().String() + " User is not found!")
		return &user_service.ChangePasswordResponse{User: mapper.MapFindUser(user)}, nil
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(requestDTO.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(requestDTO.Password)
	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		fmt.Println(time.Now().String() + " Password doesn't in valid format!")

		return &user_service.ChangePasswordResponse{User: mapper.MapFindUser(user)}, nil
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		fmt.Println(time.Now().String() + " Failed changing password!")

		return &user_service.ChangePasswordResponse{User: mapper.MapFindUser(user)}, nil
	}

	return &user_service.ChangePasswordResponse{User: mapper.MapFindUser(user)}, nil
}
