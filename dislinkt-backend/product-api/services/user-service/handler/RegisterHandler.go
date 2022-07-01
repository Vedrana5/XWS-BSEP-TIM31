package handler

import (
	"common/module/proto/user_service"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"log"
	"net/smtp"
	"time"
	"user/module/dto"
	"user/module/mapper"
	"user/module/model"
	"user/module/service"
	"user/module/util"
)

type RegisterHandler struct {
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

func NewRegisterHandler(confirmationTokenService *service.ConfirmationTokenService, passwordUtil *util.PasswordUtil, LogInfo *logrus.Logger, LogError *logrus.Logger, userService *service.UserService) *RegisterHandler {
	return &RegisterHandler{confirmationTokenService, passwordUtil, userService, LogInfo, LogError}
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
