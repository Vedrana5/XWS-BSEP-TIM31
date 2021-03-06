package handler

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
	"user/module/dto"
	"user/module/service"
	"user/module/util"
)

type LogInHandler struct {
	ValidationCodeService *service.ValidationCodeService
	UserService           *service.UserService
	PasswordUtil          *util.PasswordUtil
	LogInfo               *logrus.Logger
	LogError              *logrus.Logger
}

func (handler *LogInHandler) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

//LogInPasswordLess
func (handler *LogInHandler) LogInPasswordless(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var logInUserPasswordlessDTO dto.LogInUserPasswordlessDTO
	if err := json.NewDecoder(r.Body).Decode(&logInUserPasswordlessDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogInPasswordLess",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to LogInUserDTO!")
		fmt.Println(time.Now().String() + " Wrong cast json to LogInUserDTO!")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.UserService.FindByUserName(logInUserPasswordlessDTO.Username)

	//token
	token, err := CreateToken(user.Username)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogInPasswordLess",
			"timestamp": time.Now().String(),
		}).Error("Failed creating AWT token!")
		fmt.Println(time.Now().String() + " Failed creating AWT token!")

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
		"location":  "LogInHandler",
		"action":    "LogInPasswordLess",
		"timestamp": time.Now().String(),
	}).Info("Successfully sign in user!")
	fmt.Println(time.Now().String() + " Successfully sign in user!")

	var validation_code = handler.ValidationCodeService.FindByCode(uuid.MustParse(logInUserPasswordlessDTO.Code))
	if validation_code == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogInPasswordLess",
			"timestamp": time.Now().String(),
		}).Error("Validation_code is null!")
		fmt.Println(time.Now().String() + " Validation_code is null!")

		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if err := handler.ValidationCodeService.UpdateValidationCodeUsing(validation_code.Code); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "LogInPasswordLess",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		fmt.Println(time.Now().String() + " Failed changing password!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

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

//GetUserIDFromJWTToken
func (handler *LogInHandler) GetUserIDFromJWTToken(w http.ResponseWriter, r *http.Request) {
	token, err := VerifyToken(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "LogInHandler",
			"action":    "GetUserIDFromJWTToken",
			"timestamp": time.Now().String(),
		}).Error("Failed verified token!")
		fmt.Println(time.Now().String() + " Failed verified token!")

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
		"location":  "LogInHandler",
		"action":    "GetUserIDFromJWTToken",
		"timestamp": time.Now().String(),
	}).Error("Token doesn't valid!")
	fmt.Println(time.Now().String() + " Token doesn't valid!")
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
