package handler

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/service"

	"gopkg.in/go-playground/validator.v9"
)

type CommentHandler struct {
	Service   *service.CommentService
	PostService *service.PostService
	Validator *validator.Validate
	/*	LogInfo   *logrus.Logger
		LogError  *logrus.Logger*/
}
