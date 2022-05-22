package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type RegisterHandler struct {
	PasswordUtil *util.PasswordUtil
	UserService  *service.UserService
	LogInfo      *logrus.Logger
	LogError     *logrus.Logger
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

	gender := model.OTHER
	switch registeredUserDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	typeofUser := model.ADMIN
	switch registeredUserDTO.TypeOfUser {
	case "REGISTERED_USER":
		typeofUser = model.REGISTERED_USER
	case "UNREGISTERED_USER":
		typeofUser = model.UNREGISTERED_USER
	}

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
	}

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
