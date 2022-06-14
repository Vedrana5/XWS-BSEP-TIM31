package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"
	"github.com/form3tech-oss/jwt-go"
	"github.com/sirupsen/logrus"
)

type LogInHandler struct {
	UserService  *service.UserService
	PasswordUtil *util.PasswordUtil
	LogInfo      *logrus.Logger
	LogError     *logrus.Logger
}

func (handler *LogInHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var logInUserDTO dto.LogInUserDTO
	if err := json.NewDecoder(r.Body).Decode(&logInUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "LOG85310",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to LogInUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.UserService.FindByUserName(logInUserDTO.Username)
	validPassword := handler.PasswordUtil.IsValidPassword(logInUserDTO.Password)
	plainPassword := ""
	if !validPassword {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "LOG85310",
			"timestamp": time.Now().String(),
		}).Error("Password isn't in valid format!")
		w.WriteHeader(http.StatusBadRequest)
		return
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
			"location":  "UserHandler",
			"action":    "LOG85310",
			"timestamp": time.Now().String(),
		}).Error("Failed sign up!")
		w.WriteHeader(http.StatusConflict)
		return
	}

	//token
	token, err := CreateToken(user.Username)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "LOG85310",
			"timestamp": time.Now().String(),
		}).Error("Failed creating AWT token!")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	logInResponse := dto.LogInResponseDTO{
		ID:         user.ID,
		Token:      token,
		TypeOfUser: user.TypeOfUser,
	}
	fmt.Print("Response je ", logInResponse)
	logInResponseJson, _ := json.Marshal(logInResponse)
	w.Write(logInResponseJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisteredUserHandler",
		"action":    "LOG85310",
		"timestamp": time.Now().String(),
	}).Info("Successfully sign up user! User: " + user.Username)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func CreateToken(userName string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userName
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (handler *LogInHandler) GetUserIDFromJWTToken(w http.ResponseWriter, r *http.Request) {
	token, err := VerifyToken(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "GetUserIDFromJWTToken",
			"timestamp": time.Now().String(),
		}).Error("Failed verified token!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := fmt.Sprintf("%s", claims["user_id"])
		retValJson, _ := json.Marshal(userId)
		w.Write(retValJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "UserHandler",
		"action":    "GetUserIDFromJWTToken",
		"timestamp": time.Now().String(),
	}).Error("Token doesn't valid!")
	w.WriteHeader(http.StatusBadRequest)
}

func getUserNameFromJWT(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := fmt.Sprintf("%s", claims["user_id"])
		return userId, nil
	}
	return "", err
}
