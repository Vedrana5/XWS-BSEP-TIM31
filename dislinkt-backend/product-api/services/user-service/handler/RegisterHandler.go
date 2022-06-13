package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RegisterHandler struct {
	ConfirmationTokenService *service.ConfirmationTokenService
	PasswordUtil             *util.PasswordUtil
	UserService              *service.UserService
	LogInfo                  *logrus.Logger
	LogError                 *logrus.Logger
}

func (handler *RegisterHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var registeredUserDTO dto.RegisteredUserDTO

	if err := json.NewDecoder(r.Body).Decode(&registeredUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to RegisteredUserDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(registeredUserDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered username!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(registeredUserDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered email!")
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
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
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
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
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Failed creating confirmation token for user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.SendConfirmationMail(registeredUser, confirmationToken.ConfirmationToken)

	if err := handler.UserService.CreateUser(&registeredUser); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Failed creating basic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *RegisterHandler) SendConfirmationMail(user model.User, token uuid.UUID) {
	// Sender data.
	from := "pera08085@gmail.com"
	password := "pericaProba"

	// Receiver email address.
	to := []string{
		user.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Dear " + user.FirstName + ",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttp://localhost:8082/confirmRegistration/" + token.String() + "/" + user.ID.String())

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
		"location":  "RegisteredUserHandler",
		"action":    "SEDCONFMAIL227",
		"timestamp": time.Now().String(),
	}).Info("Successfully sended email with confirmation token!")

}
